package home

import (
	"HoneyOrangeHelper/internal/config"
	"HoneyOrangeHelper/internal/global"
	"HoneyOrangeHelper/internal/helper_cmd"
	process "HoneyOrangeHelper/internal/utils"
	"HoneyOrangeHelper/pages/server"
	"HoneyOrangeHelper/pkg/utils"
	"context"
	"slices"
	"strconv"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/issueye/ipc_grpc/grpc/pb"
	"github.com/ying32/govcl/vcl"
)

// 初始化数据
func (f *TFrmHome) InitData() {
	f.ActiveItem = 0
	f.SrvList = make([]*SrvObject, 0)
	f.IMList = make([]*ItemMenuObject, 0)

	list, err := config.GetServerList("")
	if err != nil {
		return
	}

	// 加载到主页面中，添加 PageControl 页面
	for _, v := range list {
		f.AppendServer(v)
	}

	if len(f.SrvList) > 0 {
		f.ActiveItem = f.SrvList[0].Info.Code
		f.SrvList[0].Frm.OnBtn_itemClick(nil)
		f.SrvList[0].Frm.Btn_item.SetImageIndex(0)
	}

}

func (f *TFrmHome) LoadToolPlugins() {
	// 加载工具插件
	tpList, err := config.GetToolPluginList("")
	if err != nil {
		return
	}

	for _, v := range tpList {
		f.AddMenuItem(v)
	}
}

// 添加菜单项
func (f *TFrmHome) AddMenuItem(tp *config.ToolPlugin) {
	menu := vcl.NewMenuItem(f)
	menu.SetCaption(tp.Name)
	f.Meu_tool_plugin_mana.Add(menu)

	// 创建对象
	imObj := &ItemMenuObject{Menu: menu, Info: tp}
	imObj.ctx, imObj.cancel = context.WithCancel(context.Background())
	// 启动拓展插件
	global.Sugared.Debugf("启动插件：%s", tp.Path)
	result, err := helper_cmd.Run(100)(imObj.ctx, true, tp.Path)
	if err != nil {
		global.Sugared.Debugf("启动插件失败：%s", err.Error())
		return
	}

	go func() {
		for msg := range result.Msg {
			global.Sugared.Named(tp.Name).Debug(msg)
		}
	}()

	imObj.RunResult = result
	f.IMList = append(f.IMList, imObj)
	menu.SetOnClick(f.MenuItemOnclick(tp))
}

func (f *TFrmHome) MenuItemOnclick(tp *config.ToolPlugin) func(sender vcl.IObject) {
	return func(sender vcl.IObject) {
		// vcl.ShowMessage(tp.Name)
		global.PluginSrv.Commands.Show(tp.CookieKey)
	}
}

// 添加服务
func (f *TFrmHome) AppendServer(cfg *config.Server) {
	Id := utils.GenID()
	cfg.Code = Id
	ctx, cancel := context.WithCancel(context.Background())
	frm := server.NewItemForm(ctx, f, Id, f.Sb_box, f.Pnl_body, cfg)
	obj := &SrvObject{Info: cfg, Frm: frm, ctx: ctx, cancel: cancel}
	obj.Id = Id
	f.SrvList = append(f.SrvList, obj)
	frm.Show()

	info := &pb.ServerInfo{
		Id:          strconv.FormatInt(cfg.Code, 10),
		Name:        cfg.Name,
		ProcessName: cfg.Path,
		State:       pb.StateType_ST_STOP,
	}
	f.PushEvent(pb.EventType_ET_ADD, info)
}

func (f *TFrmHome) PushEvent(t pb.EventType, info *pb.ServerInfo) {
	list, err := config.GetToolPluginList("")
	if err != nil {
		return
	}

	for _, v := range list {
		global.PluginSrv.Events.PushEvent(v.CookieKey, t, info)
	}
}

func (f *TFrmHome) SrvHelper(msgInfo *message.Message) error {
	defer msgInfo.Ack()

	return nil
}

func (f *TFrmHome) RefreshToolPlugin(msgInfo *message.Message) error {
	defer msgInfo.Ack()

	// 清空菜单
	p := new(process.Process)
	for _, item := range f.IMList {
		// 关闭进程
		// 判断是否有进程
		if item.RunResult.Pid > 0 {
			err := p.KillProcessAndChildren(item.RunResult.Pid)
			if err != nil {
				return err
			}
		}

		// 释放内存
		item.Menu.Free()

		// 取消监听
		item.cancel()
	}

	// 清空菜单对象
	f.IMList = make([]*ItemMenuObject, 0)

	// 重新加载菜单
	tpList, err := config.GetToolPluginList("")
	if err != nil {
		return err
	}

	for _, v := range tpList {
		f.AddMenuItem(v)
	}

	return nil
}

func (f *TFrmHome) EventMonitor(ctx context.Context) {
	// 移除服务事件
	removeMessages, err := global.PubSub.Subscribe(context.Background(), global.TOPIC_SERVER_REMOVE)
	if err != nil {
		return
	}

	// 选择服务事件
	indexMessages, err := global.PubSub.Subscribe(context.Background(), global.TOPIC_SERVER_INDEX)
	if err != nil {
		return
	}

	// 刷新插件事件
	tpRefreshMessages, err := global.PubSub.Subscribe(context.Background(), global.TOPIC_SERVER_REFRESH_TOOL_PLUGIN)
	if err != nil {
		return
	}

	// 帮助事件
	srvHelperMessages, err := global.PubSub.Subscribe(context.Background(), global.TOPIC_SERVER_HELPER)
	if err != nil {
		return
	}

	go func(c context.Context) {
		for {
			select {
			case msg := <-removeMessages:
				{
					err = f.closeItem(msg)
					if err != nil {
						global.Sugared.Errorf("关闭项目失败 %s", err.Error())
						continue
					}
				}
			case msg := <-indexMessages:
				{
					err = f.indexItem(msg)
					if err != nil {
						global.Sugared.Errorf("选择项目失败 %s", err.Error())
						continue
					}
				}
			case msg := <-tpRefreshMessages:
				{
					err = f.RefreshToolPlugin(msg)
					if err != nil {
						global.Sugared.Errorf("刷新插件失败 %s", err.Error())
						continue
					}
				}
			case msg := <-srvHelperMessages:
				{
					err = f.SrvHelper(msg)
					if err != nil {
						global.Sugared.Errorf("帮助失败 %s", err.Error())
						continue
					}
				}
			case <-c.Done():
				return
			}
		}
	}(ctx)
}

// 选中服务
func (f *TFrmHome) indexItem(msgInfo *message.Message) error {
	defer msgInfo.Ack()

	data := string(msgInfo.Payload)
	cfg, err := config.Server{}.FromToJson(data)
	if err != nil {
		return err
	}

	for _, item := range f.SrvList {
		if cfg.Code == item.Info.Code {
			item.Frm.Btn_item.SetImageIndex(0)
		} else {
			item.Frm.Btn_item.SetImageIndex(-1)
		}
	}

	return nil
}

// 关闭服务
func (f *TFrmHome) closeItem(msgInfo *message.Message) error {
	defer msgInfo.Ack()

	data := string(msgInfo.Payload)
	cfg, err := config.Server{}.FromToJson(data)
	if err != nil {
		return err
	}

	f.SrvList = slices.DeleteFunc(f.SrvList, func(v *SrvObject) bool {
		if v.Id == cfg.Code {
			v.Frm.CloseWindow()
			return true
		}

		return false
	})

	info := &pb.ServerInfo{
		Id:          strconv.FormatInt(cfg.Code, 10),
		Name:        cfg.Name,
		ProcessName: cfg.Path,
		State:       pb.StateType_ST_STOP,
	}
	f.PushEvent(pb.EventType_ET_REMOVE, info)

	return nil
}
