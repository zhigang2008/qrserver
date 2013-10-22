
// ServerMgrDlg.cpp : 实现文件
//

#include "stdafx.h"
#include "ServerMgr.h"
#include "ServerMgrDlg.h"
//#include "afxdialogex.h"

#ifdef _DEBUG
#define new DEBUG_NEW
#endif

const UINT WM_TASKBARCREATED = ::RegisterWindowMessage(_T("TaskbarCreated"));


// 用于应用程序“关于”菜单项的 CAboutDlg 对话框

class CAboutDlg : public CDialogEx
{
public:
	CAboutDlg();

// 对话框数据
	enum { IDD = IDD_ABOUTBOX };

	protected:
	virtual void DoDataExchange(CDataExchange* pDX);    // DDX/DDV 支持

// 实现
protected:
	DECLARE_MESSAGE_MAP()
};

CAboutDlg::CAboutDlg() : CDialogEx(CAboutDlg::IDD)
{
}

void CAboutDlg::DoDataExchange(CDataExchange* pDX)
{
	CDialogEx::DoDataExchange(pDX);
}

BEGIN_MESSAGE_MAP(CAboutDlg, CDialogEx)
	
END_MESSAGE_MAP()


// CServerMgrDlg 对话框

#define WM_NOTIFYICON (WM_USER +101)

CServerMgrDlg::CServerMgrDlg(CWnd* pParent /*=NULL*/)
	: CDialogEx(CServerMgrDlg::IDD, pParent)
	, pMenuContext(NULL)
{
	m_hIcon = AfxGetApp()->LoadIcon(IDR_MAINFRAME);
}

void CServerMgrDlg::DoDataExchange(CDataExchange* pDX)
{
	CDialogEx::DoDataExchange(pDX);
}

BEGIN_MESSAGE_MAP(CServerMgrDlg, CDialogEx)
	ON_WM_SYSCOMMAND()
	ON_WM_PAINT()
	ON_WM_QUERYDRAGICON()
	//add
	ON_MESSAGE (WM_NOTIFYICON ,&OnNotifyIcon)
	ON_WM_CLOSE()
	ON_REGISTERED_MESSAGE(WM_TASKBARCREATED,OnTaskBarCreated)
	ON_COMMAND(ID_MENU_APP_EXIT, &CServerMgrDlg::OnAppExit)
	ON_BN_CLICKED(IDC_BUTTON11, &CServerMgrDlg::OnBnClickedButton11)
	ON_BN_CLICKED(IDC_BUTTON12, &CServerMgrDlg::OnBnClickedButton12)
	ON_BN_CLICKED(IDC_BUTTON13, &CServerMgrDlg::OnBnClickedButton13)
	ON_BN_CLICKED(IDC_BUTTON21, &CServerMgrDlg::OnBnClickedButton21)
	ON_BN_CLICKED(IDC_BUTTON22, &CServerMgrDlg::OnBnClickedButton22)
	ON_BN_CLICKED(IDC_BUTTON23, &CServerMgrDlg::OnBnClickedButton23)
	ON_COMMAND(ID_MENU_SERVER_START, &CServerMgrDlg::OnMenuServerStart)
	ON_COMMAND(ID_MENU_SERVER_STOP, &CServerMgrDlg::OnMenuServerStop)
	ON_COMMAND(ID_MENU_DATA_START, &CServerMgrDlg::OnMenuDataStart)
	ON_COMMAND(ID_MENU_DATA_STOP, &CServerMgrDlg::OnMenuDataStop)
END_MESSAGE_MAP()


#define WM_NOTIFYICON (WM_USER +101)
// CServerMgrDlg 消息处理程序

BOOL CServerMgrDlg::OnInitDialog()
{
	CDialogEx::OnInitDialog();

	// 将“关于...”菜单项添加到系统菜单中。

	// IDM_ABOUTBOX 必须在系统命令范围内。
	ASSERT((IDM_ABOUTBOX & 0xFFF0) == IDM_ABOUTBOX);
	ASSERT(IDM_ABOUTBOX < 0xF000);

	CMenu* pSysMenu = GetSystemMenu(FALSE);
	if (pSysMenu != NULL)
	{
		BOOL bNameValid;
		CString strAboutMenu;
		bNameValid = strAboutMenu.LoadString(IDS_ABOUTBOX);
		ASSERT(bNameValid);
		if (!strAboutMenu.IsEmpty())
		{
			pSysMenu->AppendMenu(MF_SEPARATOR);
			pSysMenu->AppendMenu(MF_STRING, IDM_ABOUTBOX, strAboutMenu);
		}
	}

	// 设置此对话框的图标。当应用程序主窗口不是对话框时，框架将自动
	//  执行此操作
	SetIcon(m_hIcon, TRUE);			// 设置大图标
	SetIcon(m_hIcon, FALSE);		// 设置小图标

	// TODO: 在此添加额外的初始化代码
	//检查注册表
	CheckRegister();
	//检查系统服务
	if(appExist==true)
	{
		CheckService();
	}

	//初始化托盘右键菜单
	menu.LoadMenu(IDR_MENU_CONTEXT);
	pMenuContext = menu.GetSubMenu(0);
	
	//初始化状态
	InitItemState();

	//添加系统托盘图标
	InstallIcons();
	 

	return TRUE;  // 除非将焦点设置到控件，否则返回 TRUE
}

void CServerMgrDlg::OnSysCommand(UINT nID, LPARAM lParam)
{
	if ((nID & 0xFFF0) == IDM_ABOUTBOX)
	{
		CAboutDlg dlgAbout;
		dlgAbout.DoModal();
	}
	else
	{
		CDialogEx::OnSysCommand(nID, lParam);
	}
}

// 如果向对话框添加最小化按钮，则需要下面的代码
//  来绘制该图标。对于使用文档/视图模型的 MFC 应用程序，
//  这将由框架自动完成。

void CServerMgrDlg::OnPaint()
{
	if (IsIconic())
	{
		CPaintDC dc(this); // 用于绘制的设备上下文

		SendMessage(WM_ICONERASEBKGND, reinterpret_cast<WPARAM>(dc.GetSafeHdc()), 0);

		// 使图标在工作区矩形中居中
		int cxIcon = GetSystemMetrics(SM_CXICON);
		int cyIcon = GetSystemMetrics(SM_CYICON);
		CRect rect;
		GetClientRect(&rect);
		int x = (rect.Width() - cxIcon + 1) / 2;
		int y = (rect.Height() - cyIcon + 1) / 2;

		// 绘制图标
		dc.DrawIcon(x, y, m_hIcon);
	}
	else
	{
		CDialogEx::OnPaint();
	}
}

//当用户拖动最小化窗口时系统调用此函数取得光标
//显示。
HCURSOR CServerMgrDlg::OnQueryDragIcon()
{
	return static_cast<HCURSOR>(m_hIcon);
}

//添加
LRESULT CServerMgrDlg::OnNotifyIcon(WPARAM wParam, LPARAM lParam)
{
	switch(lParam)//根据lParam判断相对应的事件
	{
	case WM_LBUTTONDBLCLK://如果左键双击托盘图标,则显示窗体
		ShowWindow(SW_SHOWNORMAL );
		SetForegroundWindow();
		break;
	case WM_RBUTTONUP://如果右键菜单弹起,则弹出菜单
		CPoint pos;
		GetCursorPos(&pos);
		if(pMenuContext != NULL)
		{
			SetForegroundWindow();
			pMenuContext->TrackPopupMenu(TPM_RIGHTBUTTON|TPM_RIGHTALIGN,pos.x+1,pos.y+1,this);
		}
		break;
	}
	return 0;
}
void CServerMgrDlg::OnClose()
{
	// TODO: 在此添加消息处理程序代码和/或调用默认值
	nid.uFlags = NIF_ICON|NIF_MESSAGE|NIF_TIP|NIF_INFO;
	wcscpy(nid.szTip, L"DQS平台管理器");
	wcscpy(nid.szInfo, L"管理器仍在运行.");
	wcscpy(nid.szInfoTitle, L"I'm here");
	Shell_NotifyIcon(NIM_MODIFY,&nid);
	ShowWindow (SW_HIDE);
	//CDialog::OnClose();
}

void CServerMgrDlg::OnMenuShow()
{
	// TODO: 在此添加命令处理程序代码
	if(IsWindowVisible ())
		ShowWindow(SW_HIDE );
	else
		ShowWindow(SW_SHOWNORMAL );
}

BOOL CServerMgrDlg::InstallIcons()
{
	//设置托盘图标
	nid.cbSize =sizeof(NOTIFYICONDATA);
	nid.hWnd =m_hWnd;
	nid.uID = 0;
	nid.hIcon =m_hIcon;
	nid.uFlags = NIF_ICON|NIF_MESSAGE|NIF_TIP;
	nid.uCallbackMessage =WM_NOTIFYICON;
	wcscpy(nid.szTip, L"DQS平台管理器");
	//wcscpy(nid.szInfo, L"管理器已运行");
	//wcscpy(nid.szInfoTitle, L"I'm here");
	nid.dwInfoFlags = NIIF_INFO;
	nid.uTimeout = 1000;
	nid.uVersion = NOTIFYICON_VERSION;
	//添加系统托盘图标
	return (Shell_NotifyIcon(NIM_ADD,&nid));
}

LRESULT CServerMgrDlg::OnTaskBarCreated(WPARAM wParam, LPARAM lParam)
{
	VERIFY(InstallIcons());
    return 0;
}

//程序退出
void CServerMgrDlg::OnAppExit()
{
	// TODO: 在此添加命令处理程序代码
	::Shell_NotifyIcon (NIM_DELETE ,&nid);
	PostMessage(WM_CLOSE);
	exit(0);
}

void CServerMgrDlg::CheckRegister()
{
	 CRegKey reg;       //定义一个CRegKey对象
     CString key("software\\DqsServer");
	 LONG result=reg.Open(HKEY_LOCAL_MACHINE,key,KEY_READ);
	 if (ERROR_SUCCESS !=result)
	 {
		// MessageBox(_T("注册表打开失败"));
		 appExist=false;
	 }else{
		 appExist=true;
			 
		 LPTSTR dir;
         DWORD pCount=1024;
         if(reg.QueryStringValue(_T("installDir"),dir,&pCount)==ERROR_SUCCESS)
	     {
				 installDir=dir;
		 }else{
			 installDir=CString("C:\\Program Files\\dqs");
		 }

		 DWORD database;
		 if(reg.QueryDWORDValue(_T("database"),database)==ERROR_SUCCESS)
		 {
		   if (database==1)
		   {
			   localDatabase=true;
		   }else{
			   localDatabase=false;
		   }
		 }else{
			 localDatabase=false;
		   }
	 }

	 reg.Close();
}




void CServerMgrDlg::OnBnClickedButton11()
{
	
	 if(StartService(sService,0,NULL))
	 {
		 statusChange_ServerRun();
		 bServerStart=true;
	 }
}


void CServerMgrDlg::OnBnClickedButton12()
{
	 if(ControlService(sService,SERVICE_CONTROL_STOP,returnstatus))
	 {
		 statusChange_ServerStop();
		 bServerStart=false;
	 }
}


void CServerMgrDlg::OnBnClickedButton13()
{
	// TODO: 在此添加控件通知处理程序代码
}


void CServerMgrDlg::OnBnClickedButton21()
{
	if(StartService(dService,0,NULL))
	 {
		 statusChange_DataRun();
		 bDataStart=true;
	 }
}


void CServerMgrDlg::OnBnClickedButton22()
{
	if(ControlService(dService,SERVICE_CONTROL_STOP,returnstatus))
	 {
		 statusChange_DataStop();
		 bDataStart=false;
	 }
}


void CServerMgrDlg::OnBnClickedButton23()
{
	// TODO: 在此添加控件通知处理程序代码
}


void CServerMgrDlg::OnMenuServerStart()
{
	// TODO: 在此添加命令处理程序代码
}


void CServerMgrDlg::OnMenuServerStop()
{
	// TODO: 在此添加命令处理程序代码
}


void CServerMgrDlg::OnMenuDataStart()
{
	// TODO: 在此添加命令处理程序代码
}


void CServerMgrDlg::OnMenuDataStop()
{
	// TODO: 在此添加命令处理程序代码
}


// 设置各控件的初始化状态
void CServerMgrDlg::InitItemState()
{
	if(appExist==false)
	{
		
		GetDlgItem(IDC_GROUP1)->EnableWindow(FALSE);
		GetDlgItem(IDC_GROUP2)->EnableWindow(FALSE);
		GetDlgItem(IDC_BUTTON11)->EnableWindow(FALSE);
		GetDlgItem(IDC_BUTTON12)->EnableWindow(FALSE);
		GetDlgItem(IDC_BUTTON13)->EnableWindow(FALSE);
		GetDlgItem(IDC_BUTTON21)->EnableWindow(FALSE);
		GetDlgItem(IDC_BUTTON22)->EnableWindow(FALSE);
		GetDlgItem(IDC_BUTTON23)->EnableWindow(FALSE);
		pMenuContext->EnableMenuItem(ID_MENU_SERVER_START,MF_GRAYED);//MF_BYCOMMAND | MF_DISABLED | MF_GRAYED
		pMenuContext->EnableMenuItem(ID_MENU_SERVER_STOP,MF_GRAYED);
		pMenuContext->EnableMenuItem(ID_MENU_DATA_START,MF_GRAYED);
		pMenuContext->EnableMenuItem(ID_MENU_DATA_STOP,MF_GRAYED);
	}else{
		
		if(bServerStart==true){
			statusChange_ServerRun();
		}else{
			statusChange_ServerStop();
		}
		if(bDataStart==true){
			statusChange_DataRun();
		}else{
			statusChange_DataStop();
		}
		
		//无本地数据库
		if(localDatabase==false)
		{
			GetDlgItem(IDC_GROUP2)->EnableWindow(FALSE);
			GetDlgItem(IDC_BUTTON21)->EnableWindow(FALSE);
		    GetDlgItem(IDC_BUTTON22)->EnableWindow(FALSE);
		    GetDlgItem(IDC_BUTTON23)->EnableWindow(FALSE);
			pMenuContext->EnableMenuItem(ID_MENU_DATA_START,MF_GRAYED);
		    pMenuContext->EnableMenuItem(ID_MENU_DATA_STOP,MF_GRAYED);
		}
	}

}


void CServerMgrDlg::CheckService()
{
	 
	 hSCM = OpenSCManager(NULL, NULL, SC_MANAGER_ALL_ACCESS);
	 sService=OpenService(hSCM,_T("DQS_Server"),SERVICE_ALL_ACCESS);
	 dService=OpenService(hSCM,_T("DQS_MongoDB"),SERVICE_ALL_ACCESS);

	 if(sService!=NULL)
	 {
		 QueryServiceStatus(sService,returnstatus);
	     if(returnstatus->dwCurrentState==SERVICE_RUNNING)
	     {
		   bServerStart=true;
	      }else{
		   bServerStart=false;
	      }
	 }else
	 {
		  bServerStart=false;
	 }

	 if(dService!=NULL)
	 {
	    QueryServiceStatus(dService,returnstatus);
	    if(returnstatus->dwCurrentState==SERVICE_RUNNING)
	    {
		   bDataStart=true;
	     }else{
		   bDataStart=false;
	   }
	 }
	 else
	 {
		 bDataStart=false;
	 }
	 //::CloseHandle(hSCM);
}


void CServerMgrDlg::statusChange_ServerRun(void)
{
	
	GetDlgItem(IDC_LABLE1)->SetWindowText(_T("Runing"));
	GetDlgItem(IDC_BUTTON11)->EnableWindow(FALSE);
	GetDlgItem(IDC_BUTTON12)->EnableWindow(TRUE);
}


void CServerMgrDlg::statusChange_ServerStop(void)
{
	GetDlgItem(IDC_LABLE1)->SetWindowText(_T("Stop"));
	GetDlgItem(IDC_BUTTON11)->EnableWindow(TRUE);
	GetDlgItem(IDC_BUTTON12)->EnableWindow(FALSE);
}


void CServerMgrDlg::statusChange_DataRun(void)
{
	GetDlgItem(IDC_LABLE2)->SetWindowText(_T("Runing"));
	GetDlgItem(IDC_BUTTON21)->EnableWindow(FALSE);
	GetDlgItem(IDC_BUTTON22)->EnableWindow(TRUE);
}


void CServerMgrDlg::statusChange_DataStop(void)
{
	GetDlgItem(IDC_LABLE2)->SetWindowText(_T("Stop"));
	SetDlgItemText(IDC_LABLE2,_T("Stop"));
	GetDlgItem(IDC_BUTTON21)->EnableWindow(TRUE);
	GetDlgItem(IDC_BUTTON22)->EnableWindow(FALSE);
}
