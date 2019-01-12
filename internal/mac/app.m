#include "app.h"
#include "_cgo_export.h"
#include "json.h"

@implementation App
+ (instancetype)current {
  static App *app = nil;

  @synchronized(self) {
    if (app == nil) {
      app = [[App alloc] init];
      NSApplication *nsapp = [NSApplication sharedApplication];
      nsapp.delegate = app;
    }
  }

  return app;
}

- (instancetype)init {
  self = [super init];
  self.handlers = [NSMutableDictionary dictionaryWithCapacity:64];

  [self handle:@"app.Run"
      withHandler:^(id in, NSString *returnID) {
        return [self run:in return:returnID];
      }];

  return self;
}

- (void)handle:(NSString *)method withHandler:(PlatformHandler)handler {
  self.handlers[method] = handler;
}

- (void) return:(NSString *)returnID
     withOutput:(id)out
       andError:(NSString *)err {

  char *creturnID = returnID != nil ? (char *)returnID.UTF8String : nil;
  char *cout = out != nil ? (char *)[JSONEncoder encode:out].UTF8String : nil;
  char *cerr = err != nil ? (char *)err.UTF8String : nil;

  platformReturn(creturnID, cout, cerr);
}

- (void)goCall:(NSString *)method withInput:(id)in {
  NSMutableDictionary *call = [[NSMutableDictionary alloc] init];
  call[@"Method"] = method;
  call[@"In"] = in;

  NSString *callStr = [JSONEncoder encode:call];
  goCall((char *)callStr.UTF8String);
}

- (void)run:(id)in return:(NSString *)returnID {
  [NSApp run];
  [self return:returnID withOutput:nil andError:nil];
}

- (void)applicationDidFinishLaunching:(NSNotification *)aNotification {
  [self goCall:@"app.OnRun" withInput:nil];
}

- (BOOL)applicationShouldHandleReopen:(NSApplication *)sender
                    hasVisibleWindows:(BOOL)flag {
  NSDictionary *in = @{
    @"HasVisibleWindows" : @(flag),
  };

  [self goCall:@"app.OnReopen" withInput:in];
  return YES;
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

void defer(NSString *returnID, dispatch_block_t block) {}