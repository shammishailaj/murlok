#ifndef window_h
#define window_h

#import <WebKit/WebKit.h>

@interface Window : NSWindowController <NSWindowDelegate, WKNavigationDelegate,
                                        WKUIDelegate, WKScriptMessageHandler>
@property NSURL *defaultURL;
@property(weak) WKWebView *webView;
@property(weak) NSTextField *loading;

+ (void) new:(NSDictionary *)in return:(NSString *)returnID;
- (void)setBackground:(NSString *)color frosted:(BOOL)frosted;
- (void)configureWebView;
- (void)configureTitleBar;
- (void)configureLoader;
@end

#endif /* window_h */
