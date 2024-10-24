package home

import (
	"HoneyOrangeHelper/internal/config"
	"HoneyOrangeHelper/internal/global"
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
	menu.SetOnClick(f.MenuItemOnclick(tp))

	f.IMList = append(f.IMList, &ItemMenuObject{Menu: menu, Info: tp})
}

func (f *TFrmHome) MenuItemOnclick(tp *config.ToolPlugin) func(sender vcl.IObject) {
	return func(sender vcl.IObject) {
		vcl.ShowMessage(tp.Name)
		global.PluginSrv.Commands.Show(strconv.FormatInt(tp.ID, 10))
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

	code := strconv.FormatInt(cfg.Code, 10)
	global.PluginSrv.Events.PushEvent(code, pb.EventType_ET_ADD, &pb.ServerInfo{
		Id:          code,
		Name:        cfg.Name,
		ProcessName: cfg.Path,
		State:       pb.StateType_ST_STOP,
	})
}

func (f *TFrmHome) SrvHelper(msgInfo *message.Message) error {
	defer msgInfo.Ack()

	return nil
}

func (f *TFrmHome) RefreshToolPlugin(msgInfo *message.Message) error {
	defer msgInfo.Ack()

	// 清空菜单
	for _, item := range f.IMList {
		item.Menu.Free()
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

	code := strconv.FormatInt(cfg.Code, 10)
	global.PluginSrv.Events.PushEvent(code, pb.EventType_ET_REMOVE, &pb.ServerInfo{
		Id:          code,
		Name:        cfg.Name,
		ProcessName: cfg.Path,
		State:       pb.StateType_ST_STOP,
	})

	return nil
}
