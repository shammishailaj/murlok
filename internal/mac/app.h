#ifndef app_h
#define app_h

#include "menubar.h"
#import <Cocoa/Cocoa.h>

typedef void (^PlatformHandler)(id, NSString *);

@interface App
    : NSObject <NSApplicationDelegate, NSUserNotificationCenterDelegate>
@property NSMutableDictionary<NSString *, PlatformHandler> *handlers;
@property NSString *localServerEndpoint;
@property NSDictionary *allowedHosts;
@property MenuBar *menuBar;

+ (instancetype)current;
- (instancetype)init;
- (void)handle:(NSString *)method withHandler:(PlatformHandler)handler;
+ (void) return:(NSString *)returnID
     withOutput:(id)out
       andError:(NSString *)err;
+ (void)goCall:(NSString *)method withInput:(id)in;
+ (void)run:(id)in return:(NSString *)returnID;
+ (void)debug:(NSString *)format, ...;
+ (void)error:(NSString *)format, ...;
+ (NSString *)name;
@end

void platformCall(char *rawcall);
void dispatch(NSString *returnID, dispatch_block_t block);

#endif /* app_h */