
// ServerMgrDlg.h : 头文件
//

#pragma once


// CServerMgrDlg 对话框
class CServerMgrDlg : public CDialogEx
{
// 构造
public:
	CServerMgrDlg(CWnd* pParent = NULL);	// 标准构造函数

// 对话框数据
	enum { IDD = IDD_SERVERMGR_DIALOG };

	protected:
	virtual void DoDataExchange(CDataExchange* pDX);	// DDX/DDV 支持


// 实现
protected:
	HICON m_hIcon;

	// 生成的消息映射函数
	virtual BOOL OnInitDialog();
	afx_msg void OnSysCommand(UINT nID, LPARAM lParam);
	afx_msg void OnPaint();
	afx_msg HCURSOR OnQueryDragIcon();
	//添加
	afx_msg LRESULT OnNotifyIcon(WPARAM wParam,LPARAM lParam);//托盘图标回调函数
	afx_msg void OnClose();
	afx_msg LRESULT OnTaskBarCreated(WPARAM wParam, LPARAM lParam);
	afx_msg void OnMenuShow();
	BOOL InstallIcons();
	afx_msg void OnAppExit();
	afx_msg void CheckRegister();
	DECLARE_MESSAGE_MAP()
private:
	// 托盘图标结构
	NOTIFYICONDATA nid;
	// 右键菜单
	CMenu *pMenuContext;
	CMenu menu;
	bool appExist;
    CString installDir;
    bool localDatabase;
public:
	afx_msg void OnBnClickedButton11();
	afx_msg void OnBnClickedButton12();
	afx_msg void OnBnClickedButton13();
	afx_msg void OnBnClickedButton21();
	afx_msg void OnBnClickedButton22();
	afx_msg void OnBnClickedButton23();
	afx_msg void OnMenuServerStart();
	afx_msg void OnMenuServerStop();
	afx_msg void OnMenuDataStart();
	afx_msg void OnMenuDataStop();
protected:
	// 设置各控件的初始化状态
	void InitItemState(void);
public:
	void CheckService(void);
private:
	SC_HANDLE hSCM;
	LPSERVICE_STATUS returnstatus;
	// 服务器系统服务
	SC_HANDLE sService;
	// 数据库系统服务
	SC_HANDLE dService;
	bool bServerStart,bDataStart;

	void statusChange_ServerRun(void);
	void statusChange_ServerStop(void);
	void statusChange_DataRun(void);
	void statusChange_DataStop(void);
};
