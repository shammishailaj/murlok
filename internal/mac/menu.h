#ifndef menu_h
#define menu_h

#import <Cocoa/Cocoa.h>

@interface NSMenuItem (MenuItemCategory)
- (void)setKeys:(NSString *)keys;
@end

#endif /* menu_h */