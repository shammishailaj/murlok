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
    NSString *textColor = in[@"TextColor"];
    NSString *titleBarColor = in[@"TitleBarColor"];

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

    if (![rawwin setFrameUsingName:url]) {
      [rawwin center];
    }

    Window *win = [[Window alloc] initWithWindow:rawwin];

    win.homeURL = [App current].defaultURL;
    if (url != nil) {
      win.homeURL = [NSURL URLWithString:url];
    }

    rawwin.delegate = win;
    win.windowFrameAutosaveName = url;

    [win setBackground:backgroundColor frosted:frostedBackground];
    [win configureLoaderWithTextColor:textColor];
    [win configureWebView];
    [win configureTitleBarWithBackgroundColor:titleBarColor];

    if (NSApp.keyWindow != nil) {
      NSRect bounds = NSApp.keyWindow.frame;
      bounds.origin.x += 24;
      bounds.origin.y -= 24;
      [rawwin setFrame:bounds display:YES];
    }

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

- (void)configureLoaderWithTextColor:(NSString *)textColor {
  NSBundle *mainBundle = [NSBundle mainBundle];
  NSString *imgpath =
      [mainBundle.resourcePath stringByAppendingString:@"/logo.png"];
  NSImage *icon = [[NSImage alloc] initByReferencingFile:imgpath];
  NSImageView *image = [NSImageView imageViewWithImage:icon];
  image.translatesAutoresizingMaskIntoConstraints = NO;

  NSTextField *progress = [NSTextField labelWithString:@"100%"];
  [progress setFont:[NSFont systemFontOfSize:21 weight:NSFontWeightThin]];
  progress.translatesAutoresizingMaskIntoConstraints = NO;

  if (textColor != nil) {
    progress.textColor =
        [NSColor colorWithCIColor:[CIColor colorWithHexString:textColor]];
  }

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
              constraintsWithVisualFormat:@"|-(>=16)-[box]-|"
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
  webView.customUserAgent = @"Murlok";
  webView.navigationDelegate = self;
  webView.UIDelegate = self;
  webView.allowsMagnification = YES;

  // Make background transparent.
  [webView setValue:@(NO) forKey:@"drawsBackground"];

  self.webView = webView;

  [webView addObserver:self
            forKeyPath:@"estimatedProgress"
               options:NSKeyValueObservingOptionNew
               context:nil];

  [self loadURL:self.homeURL];
}

- (void)loadURL:(NSURL *)url {
  NSURLRequest *request = [NSURLRequest requestWithURL:url];
  [self.webView loadRequest:request];
}

- (void)loadHome {
  [self loadURL:self.homeURL];
}

- (void)webView:(WKWebView *)webView
    didStartProvisionalNavigation:(WKNavigation *)navigation {
  self.err = nil;
  [self.loader setHidden:NO];
}

- (void)webView:(WKWebView *)webView
    didFinishNavigation:(WKNavigation *)navigation {
  [self.loader setHidden:YES];
  self.webView.alphaValue = 1.0;

  [self.webView evaluateJavaScript:[App current].bridgeJS
                 completionHandler:^(id res, NSError *error) {
                   if (error != nil) {
                     [App error:@"setting up murlok failed: %@",
                                error.localizedDescription];
                   }
                 }];
}

- (void)webView:(WKWebView *)webView
    didFailNavigation:(WKNavigation *)navigation
            withError:(NSError *)error {
  [App error:error.localizedDescription];
}

- (void)webView:(WKWebView *)webView
    didFailProvisionalNavigation:(WKNavigation *)navigation
                       withError:(NSError *)error {
  self.err = error;
  self.webView.alphaValue = 0.0;
  [App error:error.localizedDescription];
  [self.loadingProgress setStringValue:error.localizedDescription];
}

- (void)zoomDefault {
  self.webView.magnification = 1.0;
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

  if ([keyPath isEqual:@"estimatedProgress"] && self.err == nil) {
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

  case WKNavigationTypeReload:
  case WKNavigationTypeBackForward:
  case WKNavigationTypeOther:
  default:
    break;
  }

  decisionHandler(WKNavigationActionPolicyAllow);
}

- (WKWebView *)webView:(WKWebView *)webView
    createWebViewWithConfiguration:(WKWebViewConfiguration *)configuration
               forNavigationAction:(WKNavigationAction *)navigationAction
                    windowFeatures:(WKWindowFeatures *)windowFeatures {

  NSURL *url = navigationAction.request.URL;
  App *app = [App current];
  id allowedHost = app.allowedHosts[url.host];

  if (allowedHost == nil) {
    [[NSWorkspace sharedWorkspace] openURL:url];
    return nil;
  }

  [App goCall:@"app.Windows.NewDefault"
      withInput:@{
        @"URL" : url.absoluteString,
      }];
  return nil;
}

- (void)webView:(WKWebView *)webView
    runJavaScriptAlertPanelWithMessage:(NSString *)message
                      initiatedByFrame:(WKFrameInfo *)frame
                     completionHandler:(void (^)(void))completionHandler {
  NSAlert *alert = [[NSAlert alloc] init];
  alert.informativeText = message;

  [alert beginSheetModalForWindow:self.window
                completionHandler:^(NSModalResponse returnCode) {
                  completionHandler();
                }];
}

- (void)webView:(WKWebView *)webView
    runJavaScriptConfirmPanelWithMessage:(NSString *)message
                        initiatedByFrame:(WKFrameInfo *)frame
                       completionHandler:
                           (void (^)(BOOL result))completionHandler {
  NSAlert *alert = [[NSAlert alloc] init];
  alert.messageText = @"Confirm";
  alert.informativeText = message;
  [alert addButtonWithTitle:@"Ok"];
  [alert addButtonWithTitle:@"Cancel"];

  [alert beginSheetModalForWindow:self.window
                completionHandler:^(NSModalResponse returnCode) {
                  completionHandler(returnCode == NSAlertFirstButtonReturn);
                }];
}

- (void)webView:(WKWebView *)webView
    runJavaScriptTextInputPanelWithPrompt:(NSString *)prompt
                              defaultText:(NSString *)defaultText
                         initiatedByFrame:(WKFrameInfo *)frame
                        completionHandler:
                            (void (^)(NSString *result))completionHandler {
  NSAlert *alert = [[NSAlert alloc] init];
  alert.messageText = prompt;
  [alert addButtonWithTitle:@"Ok"];
  [alert addButtonWithTitle:@"Cancel"];

  NSTextField *input =
      [[NSTextField alloc] initWithFrame:NSMakeRect(0, 0, 294, 24)];
  input.stringValue = defaultText;
  [alert setAccessoryView:input];

  [alert beginSheetModalForWindow:self.window
                completionHandler:^(NSModalResponse returnCode) {
                  if (returnCode != NSAlertFirstButtonReturn) {
                    completionHandler(nil);
                    return;
                  }

                  [input validateEditing];
                  completionHandler(input.stringValue);
                }];
}

- (void)userContentController:(WKUserContentController *)userContentController
      didReceiveScriptMessage:(WKScriptMessage *)message {
  if (![message.name isEqual:@"murlok"]) {
    return;
  }
}

- (BOOL)windowShouldClose:(NSWindow *)sender {
  self.window = nil;
  return YES;
}

- (void)configureTitleBarWithBackgroundColor:(NSString *)color {
  self.window.titleVisibility = NSWindowTitleHidden;
  self.window.titlebarAppearsTransparent = true;

  WKWebView *webView = self.webView;
  webView.translatesAutoresizingMaskIntoConstraints = NO;

  TitleBar *titleBar = [[TitleBar alloc] init];
  titleBar.translatesAutoresizingMaskIntoConstraints = NO;

  [self.window.contentView addSubview:webView];
  [self.window.contentView
      addConstraints:
          [NSLayoutConstraint
              constraintsWithVisualFormat:@"|[webView]|"
                                  options:0
                                  metrics:nil
                                    views:NSDictionaryOfVariableBindings(
                                              webView)]];

  [self.window.contentView addSubview:titleBar];
  [self.window.contentView
      addConstraints:
          [NSLayoutConstraint
              constraintsWithVisualFormat:@"|[titleBar]|"
                                  options:0
                                  metrics:nil
                                    views:NSDictionaryOfVariableBindings(
                                              titleBar)]];

  if (color != nil) {
    titleBar.backgroundColor = color;

    [self.window.contentView
        addConstraints:
            [NSLayoutConstraint
                constraintsWithVisualFormat:@"V:|[titleBar(==22)][webView]|"
                                    options:0
                                    metrics:nil
                                      views:NSDictionaryOfVariableBindings(
                                                titleBar, webView)]];
    return;
  }

  [self.window.contentView
      addConstraints:
          [NSLayoutConstraint
              constraintsWithVisualFormat:@"V:|[titleBar(==22)]"
                                  options:0
                                  metrics:nil
                                    views:NSDictionaryOfVariableBindings(
                                              titleBar)]];
  [self.window.contentView
      addConstraints:
          [NSLayoutConstraint
              constraintsWithVisualFormat:@"V:|[webView]|"
                                  options:0
                                  metrics:nil
                                    views:NSDictionaryOfVariableBindings(
                                              webView)]];
}
@end
