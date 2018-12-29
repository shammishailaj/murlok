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

- (void)handle:(NSString *)method withHandler:(MacRPCHandler)handler {
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

- (void)run:(id)in return:(NSString *)returnID {
  [NSApp run];
  [self return:returnID withOutput:nil andError:nil];
}
@end

void platformCall(char *rawcall) {
  NSDictionary *call =
      [JSONDecoder decode:[NSString stringWithUTF8String:rawCall]];

  NSString *method = call[@"Method"];
  id in = call[@"In"];
  NSString *returnID = call[@"ReturnID"];

  @try {
    App *app = [app current];
    PlatformHandler handler = app.handlers[method];

    if (handler == nil) {
      [NSException raise:@"not handled" format:@"%@ is not handled", method];
    }

    handler(in, returnID);
  } @catch (NSException *exception) {
    NSString *err = exception.reason;
    platformReturn((char *)returnID.UTF8String, nil, (char *)err.UTF8String);
  }
}

void defer(NSString *returnID, dispatch_block_t block) {}