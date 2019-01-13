#ifndef window_h
#define window_h

#import <WebKit/WebKit.h>

@interface Window : NSWindowController <NSWindowDelegate, WKNavigationDelegate,
                                        WKUIDelegate, WKScriptMessageHandler>
@property NSURL *defaultURL;
@property(weak) WKWebView *webView;

+ (void) new:(NSDictionary *)in return:(NSString *)returnID;
- (void)setBackground:(NSString *)color frosted:(BOOL)frosted;
- (void)configureWebView;
- (void)configureTitleBar;
// - (void)configWebview;
// - (void)configTitlebar:(NSString *)title hidden:(BOOL)isHidden;
// + (void)evalJS:(NSDictionary *)in return:(NSString *)returnID;
// + (void)close:(NSDictionary *)in return:(NSString *)returnID;
@end

#endif /* window_h */
