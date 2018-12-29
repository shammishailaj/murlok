#ifndef cocoa_h
#define cocoa_h

#import <Cocoa/Cocoa.h>

void platformCall(char *rawcall);
void defer(NSString *returnID, dispatch_block_t block);

#endif /* cocoa_h */