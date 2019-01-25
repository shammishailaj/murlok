#include "window.h"
#include "app.h"
#include "color.h"
#include "image.h"
#include "titlebar.h"

@implementation Window
+ (void) new:(NSDictionary *)in return:(NSString *)returnID {
  dispatch(returnID, ^{
    NSString *url = in[@"URL"];
    NSString *backgroundColor = in[@"BackgroundColor"];
    BOOL frostedBackground = [in[@"FrostedBackground"] boolValue];

    NSWindow *rawwin = [[NSWindow alloc]
        initWithContentRect:NSMakeRect(0, 0, 1280, 720)
                  styleMask:NSWindowStyleMaskTitled |
                            NSWindowStyleMaskFullSizeContentView |
                            NSWindowStyleMaskClosable |
                            NSWindowStyleMaskMiniaturizable |
                            NSWindowStyleMaskResizable
                    backing:NSBackingStoreBuffered
                      defer:NO];
    rawwin.minSize = NSMakeSize(360, 360);

    Window *win = [[Window alloc] initWithWindow:rawwin];
    win.defaultURL = [NSURL URLWithString:url];
    rawwin.delegate = win;
    win.windowFrameAutosaveName = url;

    [win setBackground:backgroundColor frosted:frostedBackground];
    [win configureLoader];
    [win configureWebView];
    [win configureTitleBar];

    [win showWindow:nil];
    [NSApp activateIgnoringOtherApps:YES];
    [App return:returnID withOutput:nil andError:nil];
  });
}

- (void)setBackground:(NSString *)color frosted:(BOOL)frosted {
  if (frosted) {
    NSVisualEffectView *visualEffectView =
        [[NSVisualEffectView alloc] initWithFrame:self.window.frame];
    visualEffectView.material = NSVisualEffectMaterialSidebar;
    visualEffectView.blendingMode = NSVisualEffectBlendingModeBehindWindow;
    visualEffectView.state = NSVisualEffectStateFollowsWindowActiveState;
    self.window.contentView = visualEffectView;
    return;
  }

  if (color.length == 0) {
    return;
  }

  self.window.backgroundColor =
      [NSColor colorWithCIColor:[CIColor colorWithHexString:color]];
}

- (void)configureLoader {
  NSBundle *mainBundle = [NSBundle mainBundle];
  NSString *imgpath =
      [mainBundle.resourcePath stringByAppendingString:@"/logo.png"];
  NSImage *icon = [[NSImage alloc] initByReferencingFile:imgpath];
  NSImageView *image = [NSImageView imageViewWithImage:icon];
  image.translatesAutoresizingMaskIntoConstraints = NO;

  NSTextField *progress = [NSTextField labelWithString:@"100%"];
  [progress setFont:[NSFont systemFontOfSize:21 weight:NSFontWeightThin]];
  progress.translatesAutoresizingMaskIntoConstraints = NO;

  NSView *box = [[NSView alloc] initWithFrame:NSMakeRect(0, 0, 0, 0)];
  box.translatesAutoresizingMaskIntoConstraints = NO;
  [box addSubview:image];
  [box addSubview:progress];

  [box addConstraints:
           [NSLayoutConstraint
               constraintsWithVisualFormat:@"V:|[image(==32)]|"
                                   options:0
                                   metrics:nil
                                     views:NSDictionaryOfVariableBindings(
                                               image)]];
  [box addConstraints:
           [NSLayoutConstraint
               constraintsWithVisualFormat:@"|[image(==32)]-[progress]|"
                                   options:0
                                   metrics:nil
                                     views:NSDictionaryOfVariableBindings(
                                               image, progress)]];
  [box addConstraints:
           [NSLayoutConstraint
               constraintsWithVisualFormat:@"V:[progress]-4-|"
                                   options:0
                                   metrics:nil
                                     views:NSDictionaryOfVariableBindings(
                                               progress)]];
  [self.window.contentView addSubview:box];

  [self.window.contentView
      addConstraints:
          [NSLayoutConstraint
              constraintsWithVisualFormat:@"[box]-|"
                                  options:NSLayoutFormatAlignAllCenterX |
                                          NSLayoutFormatAlignAllCenterY
                                  metrics:nil
                                    views:NSDictionaryOfVariableBindings(box)]];
  [self.window.contentView
      addConstraints:
          [NSLayoutConstraint
              constraintsWithVisualFormat:@"V:[box]-|"
                                  options:NSLayoutFormatAlignAllCenterX |
                                          NSLayoutFormatAlignAllCenterY
                                  metrics:nil
                                    views:NSDictionaryOfVariableBindings(box)]];

  self.loader = box;
  self.loadingProgress = progress;
}

- (void)configureWebView {
  WKUserContentController *userContentController =
      [[WKUserContentController alloc] init];
  [userContentController addScriptMessageHandler:self name:@"murlok"];

  WKWebViewConfiguration *conf = [[WKWebViewConfiguration alloc] init];
  conf.userContentController = userContentController;
  conf.preferences.javaScriptCanOpenWindowsAutomatically = NO;
  conf.websiteDataStore = [WKWebsiteDataStore defaultDataStore];

  WKWebView *webView = [[WKWebView alloc] initWithFrame:NSMakeRect(0, 0, 0, 0)
                                          configuration:conf];
  webView.translatesAutoresizingMaskIntoConstraints = NO;
  webView.navigationDelegate = self;
  webView.UIDelegate = self;
  webView.allowsMagnification = YES;

  // Make background transparent.
  [webView setValue:@(NO) forKey:@"drawsBackground"];

  [self.window.contentView addSubview:webView];
  [self.window.contentView
      addConstraints:
          [NSLayoutConstraint
              constraintsWithVisualFormat:@"|[webView]|"
                                  options:0
                                  metrics:nil
                                    views:NSDictionaryOfVariableBindings(
                                              webView)]];
  [self.window.contentView
      addConstraints:
          [NSLayoutConstraint
              constraintsWithVisualFormat:@"V:|[webView]|"
                                  options:0
                                  metrics:nil
                                    views:NSDictionaryOfVariableBindings(
                                              webView)]];
  self.webView = webView;

  [webView addObserver:self
            forKeyPath:@"estimatedProgress"
               options:NSKeyValueObservingOptionNew
               context:nil];

  [self loadDefaultURL];
}

- (void)loadDefaultURL {
  NSURLRequest *request = [NSURLRequest requestWithURL:self.defaultURL];
  [self.webView loadRequest:request];
}

- (void)zoomDefault {
  self.webView.magnification = 1;
}

- (void)zoomIn {
  self.webView.magnification += 0.25;
}

- (void)zoomOut {
  self.webView.magnification -= 0.25;
}

- (BOOL)validateUserInterfaceItem:(id<NSValidatedUserInterfaceItem>)item {
  SEL theAction = [item action];

  if (theAction == @selector(zoomIn)) {
    return (self.webView.magnification < 4);
  }

  if (theAction == @selector(zoomOut)) {
    return (self.webView.magnification > 1);
  }

  return YES;
}

- (void)observeValueForKeyPath:(NSString *)keyPath
                      ofObject:(id)object
                        change:(NSDictionary *)change
                       context:(void *)context {

  if ([keyPath isEqual:@"estimatedProgress"]) {
    [self.loadingProgress
        setStringValue:[NSString
                           stringWithFormat:@"%.0f%%",
                                            self.webView.estimatedProgress *
                                                100]];
  }
}

- (void)webView:(WKWebView *)webView
    decidePolicyForNavigationAction:(WKNavigationAction *)navigationAction
                    decisionHandler:
                        (void (^)(WKNavigationActionPolicy))decisionHandler {
  NSURL *url = navigationAction.request.URL;
  [App debug:@"navigating to %@", url];

  switch (navigationAction.navigationType) {
  case WKNavigationTypeReload:
  case WKNavigationTypeLinkActivated:
  case WKNavigationTypeFormSubmitted:
  case WKNavigationTypeFormResubmitted: {
    App *app = [App current];
    id allowedHost = app.allowedHosts[url.host];

    if (allowedHost == nil) {
      [[NSWorkspace sharedWorkspace] openURL:url];
      decisionHandler(WKNavigationActionPolicyCancel);
      return;
    }

    break;
  }

  case WKNavigationTypeBackForward:
  case WKNavigationTypeOther:
  default:
    break;
  }

  decisionHandler(WKNavigationActionPolicyAllow);
}

- (void)webView:(WKWebView *)webView
    didStartProvisionalNavigation:(WKNavigation *)navigation {
  [self.loader setHidden:NO];
}

- (void)webView:(WKWebView *)webView
    didFinishNavigation:(WKNavigation *)navigation {
  [self.loader setHidden:YES];
}

- (void)userContentController:(WKUserContentController *)userContentController
      didReceiveScriptMessage:(WKScriptMessage *)message {
  if (![message.name isEqual:@"murlok"]) {
    return;
  }
}

- (void)configureTitleBar {
  self.window.titleVisibility = NSWindowTitleHidden;
  self.window.titlebarAppearsTransparent = true;

  TitleBar *titleBar = [[TitleBar alloc] init];
  titleBar.translatesAutoresizingMaskIntoConstraints = NO;

  [self.window.contentView addSubview:titleBar];
  [self.window.contentView
      addConstraints:
          [NSLayoutConstraint
              constraintsWithVisualFormat:@"|[titleBar]|"
                                  options:0
                                  metrics:nil
                                    views:NSDictionaryOfVariableBindings(
                                              titleBar)]];
  [self.window.contentView
      addConstraints:
          [NSLayoutConstraint
              constraintsWithVisualFormat:@"V:|[titleBar(==22)]"
                                  options:0
                                  metrics:nil
                                    views:NSDictionaryOfVariableBindings(
                                              titleBar)]];
}
@end
