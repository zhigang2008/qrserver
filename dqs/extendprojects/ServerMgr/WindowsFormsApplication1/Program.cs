using System;
//using System.Collections.Generic;
using System.Linq;
using System.Windows.Forms;
//using System.Reflection;
//using System.Runtime.InteropServices;
//using System.Diagnostics;
using System.Threading; 


namespace WindowsFormsApplication1
{
    static class Program
    {

        public static EventWaitHandle ProgramStarted;
        /// <summary>
        /// 应用程序的主入口点。
        /// </summary>
        [STAThread]
        static void Main()
        {
            // 尝试创建一个命名事件  
            bool createNew;
            ProgramStarted = new EventWaitHandle(false, EventResetMode.AutoReset, "DQS_Server_Instance", out createNew);

            // 如果该命名事件已经存在(存在有前一个运行实例)，则发事件通知并退出  
            if (!createNew)
            {
                ProgramStarted.Set();
                return;
            }

            Application.EnableVisualStyles();
            Application.SetCompatibleTextRenderingDefault(false);
            Application.Run(new frameMain());
        }

    }
}
