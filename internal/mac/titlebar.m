#include "titlebar.h"
#include "color.h"
#include "window.h"

@implementation TitleBar
- (void)mouseDragged:(nonnull NSEvent *)theEvent {
  [self.window performWindowDragWithEvent:theEvent];
}

- (void)mouseUp:(NSEvent *)event {
  Window *win = (Window *)self.window.windowController;
  [win.webView mouseUp:event];

  if (event.clickCount == 2) {
    [win.window zoom:self];
  }
}

- (void)drawRect:(NSRect)dirtyRect {
  if (self.backgroundColor != nil) {
    NSColor *color = [NSColor
        colorWithCIColor:[CIColor colorWithHexString:self.backgroundColor]];
    [color setFill];
    NSRectFill(dirtyRect);
  }

  [super drawRect:dirtyRect];
}
@end