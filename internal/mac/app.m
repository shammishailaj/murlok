#include "app.h"
#include "_cgo_export.h"
#include "json.h"
#include "window.h"

@implementation App
+ (instancetype)current {
  static App *app = nil;

  @synchronized(self) {
    if (app == nil) {
      NSApplication *nsapp = [NSApplication sharedApplication];
      app = [[App alloc] init];
      nsapp.delegate = app;
    }
  }

  return app;
}

- (instancetype)init {
  self = [super init];

  self.menuBar = [[MenuBar alloc] init];
  [self.menuBar mount];

  self.handlers = [NSMutableDictionary dictionaryWithCapacity:64];

  [self handle:@"app.Run"
      withHandler:^(id in, NSString *returnID) {
        return [App run:in return:returnID];
      }];

  [self handle:@"app.WebDir"
      withHandler:^(id in, NSString *returnID) {
        return [App webDir:in return:returnID];
      }];

  [self handle:@"windows.New"
      withHandler:^(id in, NSString *returnID) {
        return [Window new:in return:returnID];
      }];

  return self;
}

- (void)handle:(NSString *)method withHandler:(PlatformHandler)handler {
  self.handlers[method] = handler;
}

+ (void) return:(NSString *)returnID
     withOutput:(id)out
       andError:(NSString *)err {

  char *creturnID = returnID != nil ? (char *)returnID.UTF8String : nil;
  char *cout = out != nil ? (char *)[JSONEncoder encode:out].UTF8String : nil;
  char *cerr = err != nil ? (char *)err.UTF8String : nil;

  platformReturn(creturnID, cout, cerr);
}

+ (void)goCall:(NSString *)method withInput:(id)in {
  NSMutableDictionary *call = [[NSMutableDictionary alloc] init];
  call[@"Method"] = method;
  call[@"In"] = in;

  NSString *callStr = [JSONEncoder encode:call];
  goCall((char *)callStr.UTF8String);
}

+ (void)run:(id)in return:(NSString *)returnID {
  App *app = [App current];
  app.allowedHosts = in[@"AllowedHosts"];
  app.bridgeJS = in[@"BridgeJS"];
  app.defaultURL = [NSURL URLWithString:in[@"DefaultURL"]];

  NSString *settingsURL = in[@"SettingsURL"];
  if (settingsURL != nil) {
    app.settingsURL = [NSURL URLWithString:settingsURL];
  }

  [NSApp run];
  [App return:returnID withOutput:nil andError:nil];
}

- (void)applicationDidFinishLaunching:(NSNotification *)aNotification {
  [App goCall:@"app.Running" withInput:nil];
}

- (BOOL)applicationShouldHandleReopen:(NSApplication *)sender
                    hasVisibleWindows:(BOOL)flag {
  NSDictionary *in = @{
    @"HasVisibleWindows" : @(flag),
  };

  [App goCall:@"app.Reopened" withInput:in];
  return YES;
}

+ (void)webDir:(id)in return:(NSString *)returnID {
  NSBundle *mainBundle = [NSBundle mainBundle];

  NSDictionary *out = @{
    @"WebDir" : mainBundle.resourcePath,
  };

  [App return:returnID withOutput:out andError:nil];
}

+ (void)debug:(NSString *)format, ... {
  va_list vl;
  va_start(vl, format);

  NSDictionary *in = @{
    @"Msg" : [[NSString alloc] initWithFormat:format arguments:vl],
  };

  va_end(vl);

  [App goCall:@"app.Debug" withInput:in];
}

+ (void)error:(NSString *)format, ... {
  va_list vl;
  va_start(vl, format);

  NSDictionary *in = @{
    @"Msg" : [[NSString alloc] initWithFormat:format arguments:vl],
  };

  va_end(vl);

  [App goCall:@"app.Error" withInput:in];
}

+ (NSString *)name {
  NSBundle *mainBundle = [NSBundle mainBundle];
  return mainBundle.infoDictionary[@"CFBundleName"];
}

- (void)loadMurlokRepo {
  [[NSWorkspace sharedWorkspace]
      openURL:
          [NSURL URLWithString:@"https://github.com/maxence-charriere/murlok"]];
}

- (void)loadSettings {
  if (self.settingsURL == nil) {
    [App error:@"no settings url to load"];
    return;
  }

  if (NSApp.keyWindow == nil) {
    [App goCall:@"app.Windows.NewDefault"
        withInput:@{
          @"URL" : self.settingsURL,
        }];
    return;
  }

  Window *win = NSApp.keyWindow.windowController;
  [win loadURL:self.settingsURL];
}

+ (void)emit:(NSString *)event withArg:(NSDictionary *)arg {
  if (NSApp.keyWindow == nil) {
    [App goCall:@"app.Windows.NewDefault"
        withInput:@{
          @"URL" : @"",
        }];

    [App emit:event withArg:arg];
    return;
  }

  Window *win = NSApp.keyWindow.windowController;
  NSString *eval = [NSString stringWithFormat:@"murlok.onEvent('%@', %@)",
                                              event, [JSONEncoder encode:arg]];

  [win.webView
      evaluateJavaScript:eval
       completionHandler:^(id res, NSError *error) {
         if (error != nil) {
           [App error:@"emitting event failed: %@", error.localizedDescription];
         }
       }];
}
@end

void platformCall(char *rawcall) {
  NSDictionary *call =
      [JSONDecoder decode:[NSString stringWithUTF8String:rawcall]];

  NSString *method = call[@"Method"];
  id in = call[@"In"];
  NSString *returnID = call[@"ReturnID"];

  @try {
    App *app = [App current];
    PlatformHandler handler = app.handlers[method];

    if (handler == nil) {
      [NSException raise:@"not implemented"
                  format:@"%@ is not implemented", method];
    }

    handler(in, returnID);
  } @catch (NSException *exception) {
    NSString *err = exception.reason;
    platformReturn((char *)returnID.UTF8String, nil, (char *)err.UTF8String);
  }
}

void dispatch(NSString *returnID, dispatch_block_t block) {
  @try {
    block();
  } @catch (NSException *exception) {
    NSString *err = exception.reason;
    platformReturn((char *)returnID.UTF8String, nil, (char *)err.UTF8String);
  }
}