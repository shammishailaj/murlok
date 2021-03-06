#ifndef window_h
#define window_h

#import <WebKit/WebKit.h>

@interface Window : NSWindowController <NSWindowDelegate, WKNavigationDelegate,
                                        WKUIDelegate, WKScriptMessageHandler>
@property NSURL *homeURL;
@property NSError *err;
@property(weak) WKWebView *webView;
@property(weak) NSView *loader;
@property(weak) NSTextField *loadingProgress;

+ (void) new:(NSDictionary *)in return:(NSString *)returnID;
- (void)setBackground:(NSString *)color frosted:(BOOL)frosted;
- (void)configureWebView;
- (void)loadURL:(NSURL *)url;
- (void)loadHome;
- (void)zoomDefault;
- (void)zoomIn;
- (void)zoomOut;
- (void)configureTitleBarWithBackgroundColor:(NSString *)color;
- (void)configureLoaderWithTextColor:(NSString *)textColor;
@end

#endif /* window_h */
