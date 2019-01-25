#include "menubar.h"
#include "app.h"
#include "menu.h"

@implementation MenuBar
- (instancetype)init {
  self = [super init];
  self.root = [[NSMenu alloc] initWithTitle:@""];
  [self initAppMenu];
  return self;
}

- (void)initAppMenu {
  NSMenuItem *about = [[NSMenuItem alloc]
      initWithTitle:[NSString stringWithFormat:@"About %@", [App name]]
             action:@selector(orderFrontStandardAboutPanel:)
      keyEquivalent:@""];

  NSMenuItem *preferences = [[NSMenuItem alloc] initWithTitle:@"Preferences..."
                                                       action:nil
                                                keyEquivalent:@""];
  [preferences setKeys:@"cmd+,"];

  NSApp.servicesMenu = [[NSMenu alloc] initWithTitle:@"Services"];
  NSMenuItem *services = [[NSMenuItem alloc] initWithTitle:@"Services"
                                                    action:nil
                                             keyEquivalent:@""];
  services.submenu = NSApp.servicesMenu;

  NSMenuItem *hide = [[NSMenuItem alloc]
      initWithTitle:[NSString stringWithFormat:@"Hide %@", [App name]]
             action:@selector(hide:)
      keyEquivalent:@""];
  [hide setKeys:@"cmd+h"];

  NSMenuItem *hideOthers =
      [[NSMenuItem alloc] initWithTitle:@"Hide Others"
                                 action:@selector(hideOtherApplications:)
                          keyEquivalent:@""];
  [hideOthers setKeys:@"cmd+alt+h"];

  NSMenuItem *showAll =
      [[NSMenuItem alloc] initWithTitle:@"Show All"
                                 action:@selector(unhideAllApplications:)
                          keyEquivalent:@""];

  NSMenuItem *quit = [[NSMenuItem alloc]
      initWithTitle:[NSString stringWithFormat:@"Quit %@", [App name]]
             action:@selector(terminate:)
      keyEquivalent:@""];
  [quit setKeys:@"cmd+q"];

  NSMenu *menu = [[NSMenu alloc] initWithTitle:@""];
  [menu addItem:about];
  [menu addItem:[NSMenuItem separatorItem]];
  [menu addItem:preferences];
  [menu addItem:[NSMenuItem separatorItem]];
  [menu addItem:services];
  [menu addItem:[NSMenuItem separatorItem]];
  [menu addItem:hide];
  [menu addItem:hideOthers];
  [menu addItem:showAll];
  [menu addItem:[NSMenuItem separatorItem]];
  [menu addItem:quit];

  self.appMenu =
      [[NSMenuItem alloc] initWithTitle:@"" action:nil keyEquivalent:@""];
  self.appMenu.submenu = menu;
}

- (void)mount {
  NSApp.mainMenu = nil;
  [self.root removeAllItems];
  [self.root addItem:self.appMenu];
  NSApp.mainMenu = self.root;
}
@end