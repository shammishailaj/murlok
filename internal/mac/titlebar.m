#include "titlebar.h"
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
@end