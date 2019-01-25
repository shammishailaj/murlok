#ifndef menubar_h
#define menubar_h

#import <Cocoa/Cocoa.h>

@interface MenuBar : NSObject
@property NSMenu *root;
@property NSArray<NSMenuItem *> *customMenus;
@property NSMenuItem *appMenu;
@property NSMenuItem *fileMenu;
@property NSMenuItem *editMenu;
@property NSMenuItem *windowMenu;
@property NSMenuItem *helpMenu;

- (instancetype)init;
- (void)initAppMenu;
- (void)initEditMenu;
- (void)initWindowMenu;
- (void)initHelpMenu;
- (void)mount;
@end

#endif /* menubar_h */
