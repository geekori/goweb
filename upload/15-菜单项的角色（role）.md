菜单的角色就是菜单的预定义动作，通过菜单对象的role属性设置。通用的角色如下：

- undo
- redo
- cut
- copy
- paste
- pasteAndMatchStyle
- selectAll
- delete
- minimize：最小化当前窗口
- close：关闭当前窗
- quit：退出应用程序
- reload：重新装载当前窗口
- forceReload：重新装载当前窗口（不考虑缓存）
- toggleDevTools：在当前窗口显示开发者工具
- toggleFullScreen：全屏显示当前窗口
- resetZoom：重新设置当前页面的尺寸为最初的尺寸
- zoomIn：将当前页面放大10%
- zoomOut：将当前页面缩小10%
- editMenu：整个”Edit“菜单，包括Undo、Copy等。
- windowMenu：整个"Window"菜单，包括Minimize、Close等。

下面的角色仅用于Mac OS X系统。

- about：显示”关于“对话框
- hide：隐藏
- hideOthers：隐藏其他应用程序.
- unhide：取消隐藏其他应用程序.
- startSpeaking：开始说话.
- stopSpeaking ：停止说话
- front：映射arrangeInFront动作
- zoom：映射performZoom动作
- toggleTabBar：显示TabBar.
- selectNextTab：选择下一个Tab
- selectPreviousTab：选择前一个Tab
- mergeAllWindows：合并所有的窗口
- moveTabToNewWindow：移动Tab到新的窗口
- window：Window的子菜单
- help：Help的子菜单
- services：Services的子菜单
- recentDocuments：Open Recent菜单的子菜单
- clearRecentDocuments：清除最近打开的文档

下面给出一个完整的演示如何使用菜单项角色。

index.js文件
```
const electron = require('electron');
const app = electron.app;
const BrowserWindow = electron.BrowserWindow;
const Menu  = electron.Menu;


function createWindow () {

    win = new BrowserWindow({file: 'index.html'});


    win.loadFile('./index.html');

    const template = [
        {
            label: '编辑',
            submenu: [
                {
                    label: '撤销',
                    role:'undo'

                },
                {
                    label: '重做',
                    role:'redo'

                },
                {
                    label: '剪切',
                    role:'cut'
                },
                {
                    label: '复制',
                    role:'copy'
                },
                {
                    label: '粘贴',
                    role:'paste'
                }
            ]
        },
        {
            label: '调试',
            submenu: [
                {
                    label: '显示调试工具',
                    role:'toggleDevTools'

                }
            ]
        }
        ,
        {
            label: '窗口',
            submenu: [
                {
                    label: '全屏显示窗口',
                    role:'toggleFullScreen'

                },
                {
                    label: '窗口放大10%',
                    role:'zoomIn'

                },
                ,
                {
                    label: '窗口缩小10%',
                    role:'zoomOut'

                }
            ]
        }
    ];
    if (process.platform == 'darwin') {

        template.unshift({
            label: 'Mac',
            submenu: [
                {
                    label: '关于',
                    role:'about'

                },
                {
                    label: '开始说话',
                    role:'startSpeaking'

                },
                {
                    label: '停止说话',
                    role:'stopSpeaking'

                }
            ]
        })
    }
    const menu = Menu.buildFromTemplate(template);
    Menu.setApplicationMenu(menu);
    win.on('closed', () => {
      console.log('closed');
      win = null;
    })

  }

app.on('ready', createWindow)

app.on('activate', () => {

    if (win === null) {
        createWindow();
    }
})
```
运行上面程序之前，要先在index.html中加一个文本输入框，用来演示文本的复制、粘贴、剪切等功能，代码如下：
```
<!DOCTYPE html>
<html>
<head>
  <!--  指定页面编码格式  -->
  <meta charset="UTF-8">
  <!--  指定页头信息 -->
  <title>菜单项角色（role）</title>
</head>
<body>
<h1>默认模板</h1>
<textarea style="width:400px;height:300px"></textarea>
</body>
</html>
```
在前面的代码中添加菜单时考虑到了操作系统的差异，如果是Mac OS X，会在开始添加一个Mac菜单，并添加Mac OS X特有的角色作为菜单项。

Mac OS X的效果

![image.png](https://upload-images.jianshu.io/upload_images/13614258-af0926725ad32614.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

在Mac OS X下，在文本输入框输入一些文本，选中这些文本，然后单击”开始说话“菜单项，Mac OS X就会将这行文本读出来，这是苹果系统内置的功能，Windows和Linux是没这个待遇的。

Windows的效果

![image.png](https://upload-images.jianshu.io/upload_images/13614258-bcb08e74c36d00f6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


读者可以使用相应的菜单项演示各种角色的功能。如”窗口放大10%“，没单击一次，会让当前页面所有的内容放大10%。