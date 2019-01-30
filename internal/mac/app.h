#ifndef app_h
#define app_h

#include "menubar.h"
#import <Cocoa/Cocoa.h>

typedef void (^PlatformHandler)(id, NSString *);

@interface App
    : NSObject <NSApplicationDelegate, NSUserNotificationCenterDelegate>
@property NSMutableDictionary<NSString *, PlatformHandler> *handlers;
@property NSString *bridgeJS;
@property NSDictionary *allowedHosts;
@property MenuBar *menuBar;
@property NSURL *defaultURL;
@property NSURL *settingsURL;

+ (instancetype)current;
- (instancetype)init;
- (void)handle:(NSString *)method withHandler:(PlatformHandler)handler;
+ (void) return:(NSString *)returnID
     withOutput:(id)out
       andError:(NSString *)err;
+ (void)goCall:(NSString *)method withInput:(id)in;
+ (void)run:(id)in return:(NSString *)returnID;
+ (void)webDir:(id)in return:(NSString *)returnID;
+ (void)debug:(NSString *)format, ...;
+ (void)error:(NSString *)format, ...;
+ (NSString *)name;
- (void)loadSettings;
- (void)loadMurlokRepo;
+ (void)emit:(NSString *)event withArg:(NSDictionary *)arg;
@end

void platformCall(char *rawcall);
void dispatch(NSString *returnID, dispatch_block_t block);

#endif /* app_h */