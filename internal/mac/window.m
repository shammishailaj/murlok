#include "window.h"
#include "app.h"
#include "color.h"
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

- (void)configureWebView {
  WKUserContentController *userContentController =
      [[WKUserContentController alloc] init];
  [userContentController addScriptMessageHandler:self name:@"murlok"];

  WKWebViewConfiguration *conf = [[WKWebViewConfiguration alloc] init];
  conf.userContentController = userContentController;

  WKWebView *webView = [[WKWebView alloc] initWithFrame:NSMakeRect(0, 0, 0, 0)
                                          configuration:conf];
  webView.translatesAutoresizingMaskIntoConstraints = NO;
  webView.navigationDelegate = self;
  webView.UIDelegate = self;

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

  NSURLRequest *request = [NSURLRequest requestWithURL:self.defaultURL];
  [webView loadRequest:request];
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

- (void)userContentController:(WKUserContentController *)userContentController
      didReceiveScriptMessage:(WKScriptMessage *)message {
  //   if (![message.name isEqual:@"golangRequest"]) {
  //     return;
  //   }

  //   Driver *driver = [Driver current];

  //   NSDictionary *in = @{
  //     @"ID" : self.ID,
  //     @"Mapping" : message.body,
  //   };

  //   [driver.goRPC call:@"windows.OnCallback" withInput:in];
}
@end
