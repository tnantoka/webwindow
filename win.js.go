package webwindow

const WinJS = `(function() {

ObjC.import('Cocoa');
ObjC.import('WebKit');

ObjC.registerSubclass({
  name: 'WWWindowDelegate',
  superclass: 'NSObject',
  protocols: ['NSWindowDelegate'],
  methods: {
    'windowWillClose:': {
      implementation: function(notification) {
        return $.NSApp.terminate(0);
      }
    }
  }
});

// Window
var win = $.NSWindow.alloc.initWithContentRectStyleMaskBackingDefer(
    $.NSMakeRect(0, 0, {{.Width}}, {{.Height}}),
    $.NSTitledWindowMask | $.NSMiniaturizableWindowMask | $.NSClosableWindowMask,
    $.NSBackingStoreBuffered,
    false
    );

win.delegate = $.WWWindowDelegate.alloc.init;
win.title = '{{.Title}}';
win.center;

win.makeKeyAndOrderFront(win);

// Menu
var mainMenu = $.NSMenu.alloc.init;
var mainItem = $.NSMenuItem.alloc.init;
mainMenu.addItem(mainItem);
$.NSApp.setMainMenu(mainMenu);

var appMenu = $.NSMenu.alloc.init;
var quitItem = $.NSMenuItem.alloc.initWithTitleActionKeyEquivalent('Quit', 'terminate:', 'q');
appMenu.addItem(quitItem);
mainItem.setSubmenu(appMenu);

var editItem = $.NSMenuItem.alloc.init;
mainMenu.addItem(editItem);
var editMenu = $.NSMenu.alloc.initWithTitle('Edit');
editItem.setSubmenu(editMenu);

var undoItem = $.NSMenuItem.alloc.initWithTitleActionKeyEquivalent('Undo', 'undo:', 'z');
var redoItem = $.NSMenuItem.alloc.initWithTitleActionKeyEquivalent('Redo', 'redo:', 'Z');
var separatorItem = $.NSMenuItem.separatorItem;
var cutItem = $.NSMenuItem.alloc.initWithTitleActionKeyEquivalent('Cut', 'cut:', 'x');
var copyItem = $.NSMenuItem.alloc.initWithTitleActionKeyEquivalent('Copy', 'copy:', 'c');
var pasteItem = $.NSMenuItem.alloc.initWithTitleActionKeyEquivalent('Paste', 'paste:', 'v');
var selectAllItem = $.NSMenuItem.alloc.initWithTitleActionKeyEquivalent('Select All', 'selectAll:', 'a');

editMenu.addItem(undoItem);
editMenu.addItem(redoItem);
editMenu.addItem(separatorItem);
editMenu.addItem(cutItem);
editMenu.addItem(copyItem);
editMenu.addItem(pasteItem);
editMenu.addItem(selectAllItem);

// WebView
var webview = $.WebView.alloc.initWithFrameFrameNameGroupName(
    win.contentView.bounds,
    'frame',
    'group'
    );

var url = $.NSURL.alloc.initWithString('http://{{.Host}}:{{.Port}}/{{.Root}}');
var req = $.NSURLRequest.requestWithURLCachePolicyTimeoutInterval(
    url,
    $.ReloadIgnoringLocalAndRemoteCacheData,
    0
)

webview.mainFrame.loadRequest(req);
win.contentView.addSubview(webview);

// App
var app = $.NSApp;
app.setActivationPolicy($.NSApplicationActivationPolicyRegular);
app.activateIgnoringOtherApps(true);

app.run();

})();`
