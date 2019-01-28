#include "menubar.h"
#include "app.h"
#include "menu.h"

@implementation MenuBar
- (instancetype)init {
  self = [super init];
  self.root = [[NSMenu alloc] initWithTitle:@""];

  [self initAppMenu];
  [self initEditMenu];
  [self initWindowMenu];
  [self initHelpMenu];
  return self;
}

- (void)initAppMenu {
  NSMenuItem *about = [[NSMenuItem alloc]
      initWithTitle:[NSString stringWithFormat:@"About %@", [App name]]
             action:@selector(orderFrontStandardAboutPanel:)
      keyEquivalent:@""];

  NSMenuItem *preferences =
      [[NSMenuItem alloc] initWithTitle:@"Preferences..."
                                 action:@selector(loadSettings)
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

- (void)initEditMenu {
  NSMenuItem *undo = [[NSMenuItem alloc] initWithTitle:@"Undo"
                                                action:@selector(undo:)
                                         keyEquivalent:@""];
  [undo setKeys:@"cmd+z"];

  NSMenuItem *redo = [[NSMenuItem alloc] initWithTitle:@"Redo"
                                                action:@selector(redo:)
                                         keyEquivalent:@""];
  [redo setKeys:@"cmd+shift+z"];

  NSMenuItem *cut = [[NSMenuItem alloc] initWithTitle:@"Cut"
                                               action:@selector(cut:)
                                        keyEquivalent:@""];
  [cut setKeys:@"cmd+x"];

  NSMenuItem *copy = [[NSMenuItem alloc] initWithTitle:@"Copy"
                                                action:@selector(copy:)
                                         keyEquivalent:@""];
  [copy setKeys:@"cmd+c"];

  NSMenuItem *paste = [[NSMenuItem alloc] initWithTitle:@"Paste"
                                                 action:@selector(paste:)
                                          keyEquivalent:@""];
  [paste setKeys:@"cmd+v"];

  NSMenuItem *pasteAndMatchStyle =
      [[NSMenuItem alloc] initWithTitle:@"Paste and Match Style"
                                 action:@selector(pasteAsPlainText:)
                          keyEquivalent:@""];
  [pasteAndMatchStyle setKeys:@"cmd+shift+v"];

  NSMenuItem *delete = [[NSMenuItem alloc] initWithTitle:@"Delete"
                                                  action:@selector(delete:)
                                           keyEquivalent:@""];

  NSMenuItem *selectAll =
      [[NSMenuItem alloc] initWithTitle:@"Select All"
                                 action:@selector(selectAll:)
                          keyEquivalent:@""];
  [selectAll setKeys:@"cmd+a"];

  NSMenu *menu = [[NSMenu alloc] initWithTitle:@"Edit"];
  [menu addItem:undo];
  [menu addItem:redo];
  [menu addItem:[NSMenuItem separatorItem]];
  [menu addItem:cut];
  [menu addItem:copy];
  [menu addItem:paste];
  [menu addItem:pasteAndMatchStyle];
  [menu addItem:delete];
  [menu addItem:selectAll];
  [menu addItem:[NSMenuItem separatorItem]];

  self.editMenu =
      [[NSMenuItem alloc] initWithTitle:@"Edit" action:nil keyEquivalent:@""];
  self.editMenu.submenu = menu;
}

- (void)initWindowMenu {
  NSMenuItem *minimize =
      [[NSMenuItem alloc] initWithTitle:@"Minimize"
                                 action:@selector(performMiniaturize:)
                          keyEquivalent:@""];
  [minimize setKeys:@"cmd+m"];

  NSMenuItem *zoom = [[NSMenuItem alloc] initWithTitle:@"Zoom"
                                                action:@selector(performZoom:)
                                         keyEquivalent:@""];

  NSMenuItem *reload = [[NSMenuItem alloc] initWithTitle:@"Reload"
                                                  action:@selector(reload:)
                                           keyEquivalent:@""];
  [reload setKeys:@"cmd+r"];

  NSMenuItem *reloadWithoutCache =
      [[NSMenuItem alloc] initWithTitle:@"Reload Without Cache"
                                 action:@selector(reloadFromOrigin:)
                          keyEquivalent:@""];
  [reloadWithoutCache setKeys:@"cmd+shift+r"];

  NSMenuItem *back = [[NSMenuItem alloc] initWithTitle:@"Back"
                                                action:@selector(goBack:)
                                         keyEquivalent:@""];
  [back setKeys:@"cmd+["];

  NSMenuItem *forward = [[NSMenuItem alloc] initWithTitle:@"Forward"
                                                   action:@selector(goForward:)
                                            keyEquivalent:@""];
  [forward setKeys:@"cmd+]"];

  NSMenuItem *home = [[NSMenuItem alloc] initWithTitle:@"Home"
                                                action:@selector(loadHome)
                                         keyEquivalent:@""];
  [home setKeys:@"cmd+shift+h"];

  NSMenuItem *actualSize =
      [[NSMenuItem alloc] initWithTitle:@"Actual Size"
                                 action:@selector(zoomDefault)
                          keyEquivalent:@""];
  [actualSize setKeys:@"cmd+0"];

  NSMenuItem *zoomIn = [[NSMenuItem alloc] initWithTitle:@"Zoom In"
                                                  action:@selector(zoomIn)
                                           keyEquivalent:@""];
  [zoomIn setKeys:@"cmd++"];

  NSMenuItem *zoomOut = [[NSMenuItem alloc] initWithTitle:@"Zoom Out"
                                                   action:@selector(zoomOut)
                                            keyEquivalent:@""];
  [zoomOut setKeys:@"cmd+-"];

  NSMenuItem *allFront =
      [[NSMenuItem alloc] initWithTitle:@"Bring All to Front"
                                 action:@selector(arrangeInFront:)
                          keyEquivalent:@""];

  NSMenuItem *close = [[NSMenuItem alloc] initWithTitle:@"Close"
                                                 action:@selector(performClose:)
                                          keyEquivalent:@""];
  [close setKeys:@"cmd+w"];

  NSMenu *menu = [[NSMenu alloc] initWithTitle:@"Window"];
  [menu addItem:minimize];
  [menu addItem:zoom];
  [menu addItem:[NSMenuItem separatorItem]];
  [menu addItem:reload];
  [menu addItem:reloadWithoutCache];
  [menu addItem:back];
  [menu addItem:forward];
  [menu addItem:home];
  [menu addItem:[NSMenuItem separatorItem]];
  [menu addItem:actualSize];
  [menu addItem:zoomIn];
  [menu addItem:zoomOut];
  [menu addItem:[NSMenuItem separatorItem]];
  [menu addItem:allFront];
  [menu addItem:close];

  NSApp.windowsMenu = menu;
  self.windowMenu =
      [[NSMenuItem alloc] initWithTitle:@"Window" action:nil
                          keyEquivalent:@""];
  self.windowMenu.submenu = menu;
}

- (void)initHelpMenu {
  NSMenuItem *help = [[NSMenuItem alloc]
      initWithTitle:[NSString stringWithFormat:@"%@ Help", [App name]]
             action:@selector(showHelp:)
      keyEquivalent:@""];
  [help setKeys:@"cmd+?"];

  NSMenuItem *murlok =
      [[NSMenuItem alloc] initWithTitle:@"Murlok"
                                 action:@selector(loadMurlokRepo)
                          keyEquivalent:@""];

  NSMenu *menu = [[NSMenu alloc] initWithTitle:@"Help"];
  [menu addItem:help];
  [menu addItem:[NSMenuItem separatorItem]];
  [menu addItem:murlok];

  self.helpMenu =
      [[NSMenuItem alloc] initWithTitle:@"Help" action:nil keyEquivalent:@""];
  self.helpMenu.submenu = menu;
}

- (void)mount {
  NSApp.mainMenu = nil;
  [self.root removeAllItems];
  [self.root addItem:self.appMenu];
  [self.root addItem:self.editMenu];
  [self.root addItem:self.windowMenu];
  [self.root addItem:self.helpMenu];
  NSApp.mainMenu = self.root;
}
@end