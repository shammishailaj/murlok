#ifndef app_h
#define app_h

#import <Cocoa/Cocoa.h>

typedef void (^PlatformHandler)(id, NSString *);

@interface App
    : NSObject <NSApplicationDelegate, NSUserNotificationCenterDelegate>
@property NSMutableDictionary<NSString *, PlatformHandler> *handlers;

+ (instancetype)current;
- (instancetype)init;
- (void)handle:(NSString *)method withHandler:(PlatformHandler)handler;
+ (void) return:(NSString *)returnID
     withOutput:(id)out
       andError:(NSString *)err;
- (void)goCall:(NSString *)method withInput:(id)in;
- (void)run:(id)in return:(NSString *)returnID;
@end

void platformCall(char *rawcall);
void dispatch(NSString *returnID, dispatch_block_t block);

#endif /* app_h */