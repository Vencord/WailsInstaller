#import "applescript.h"

void runScript(const char *s) {
    NSString *script = [NSString stringWithUTF8String:s];
    NSAppleScript *applescript = [[NSAppleScript alloc] initWithSource:script];
    NSDictionary *errInfo = nil;
    [script executeAndReturnError:&errInfo];
}
