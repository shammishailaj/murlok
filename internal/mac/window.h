#ifndef window_h
#define window_h

#import <WebKit/WebKit.h>

@interface Window : NSWindowController <NSWindowDelegate, WKNavigationDelegate,
                                        WKUIDelegate, WKScriptMessageHandler>
@property NSURL *defaultURL;
@property NSError *err;
@property(weak) WKWebView *webView;
@property(weak) NSView *loader;
@property(weak) NSTextField *loadingProgress;

+ (void) new:(NSDictionary *)in return:(NSString *)returnID;
- (void)setBackground:(NSString *)color frosted:(BOOL)frosted;
- (void)configureWebView;
- (void)loadDefaultURL;
- (void)zoomDefault;
- (void)zoomIn;
- (void)zoomOut;
- (void)configureTitleBar;
- (void)configureLoader;
@end

#endif /* window_h */
