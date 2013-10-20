using System;
using System.Collections.Generic;
using System.Linq;
using System.Windows.Forms;
using System.Reflection;
using System.Runtime.InteropServices;
using System.Diagnostics;


namespace WindowsFormsApplication1
{
    static class Program
    {
        /// <summary>
        /// 应用程序的主入口点。
        /// </summary>
        [STAThread]
        static void Main()
        {
            Process instance = RunningInstance();
            if (instance == null)
            {
                Application.EnableVisualStyles();
                Application.SetCompatibleTextRenderingDefault(false);
                Application.Run(new frameMain());
            }
            else
            {
                HandleRunningInstance(instance);
            }
            
        }

         #region  确保程序只运行一个实例
         private static Process RunningInstance()
         {
             Process current = Process.GetCurrentProcess();
             Process[] processes = Process.GetProcessesByName(current.ProcessName);
             //遍历与当前进程名称相同的进程列表  
             foreach (Process process in processes)
             {
                 //如果实例已经存在则忽略当前进程  
                 if (process.Id != current.Id)
                 {
                     //保证要打开的进程同已经存在的进程来自同一文件路径
                     if (Assembly.GetExecutingAssembly().Location.Replace("/", "\\") == current.MainModule.FileName)
                     {
                         //返回已经存在的进程
                         return process;
                         
                     }
                 }
             }
             return null;
         }
 
         private static void HandleRunningInstance(Process instance)
         {
             MessageBox.Show("已经在运行！", instance.MainWindowHandle.ToString(), MessageBoxButtons.OK, MessageBoxIcon.Information);
             IntPtr hWnd = instance.MainWindowHandle;
             SendMessage(instance.MainWindowHandle, WM_SHOWWINDOW, 1, 0);
             ShowWindowAsync(hWnd, 1);  //调用api函数，正常显示窗口
             SetForegroundWindow(hWnd); //将窗口放置最前端
         }

         [DllImport("User32.dll")]
         private static extern bool ShowWindowAsync(System.IntPtr hWnd, int cmdShow);
         [DllImport("User32.dll")]
         private static extern bool SetForegroundWindow(System.IntPtr hWnd);
         [DllImport("user32.dll", EntryPoint = "SendMessage")]
         private static extern int SendMessage(IntPtr hwnd, int wMsg, int wParam, int lParam);
         [DllImport("user32.dll")]
         private static extern bool IsIconic(IntPtr hWnd);

         private const int WM_SHOWWINDOW = 24;
         private const int SW_SHOWNORMAL = 10;

         #endregion
    }
}
