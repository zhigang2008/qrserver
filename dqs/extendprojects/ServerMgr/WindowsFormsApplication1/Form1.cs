using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;
using System.ServiceProcess;
using Microsoft.Win32;
using System.Diagnostics;
using System.Threading;

namespace WindowsFormsApplication1
{
    public partial class frameMain : Form
    {
        bool appExist = false;
        ServiceController sc_server, sc_data;
        //是否安装本地数据库
        string localDatabase = "0"; 
        //安装路径
        string installDir="";

        public frameMain()
        {
            InitializeComponent();
            // 单实例实现
            ThreadPool.RegisterWaitForSingleObject(Program.ProgramStarted, OnProgramStarted, null, -1, false);
        }

        private void frameMain_FormClosing(object sender, FormClosingEventArgs e)
        {

            e.Cancel = true; // 取消关闭窗体

            this.Visible=false;
            this.ShowInTaskbar = false;
            this.mgrNotifyIcon.Visible = true;//显示托盘图标
        }

        private void frameMain_Resize(object sender, EventArgs e)
        {
            if (this.WindowState == FormWindowState.Minimized)
            {
                this.WindowState = FormWindowState.Minimized;
                this.Visible = false;
                this.mgrNotifyIcon.Visible = true;

            }
        }


        private void mgrNotifyIcon_MouseDoubleClick(object sender, MouseEventArgs e)
        {
            if (this.WindowState == FormWindowState.Minimized || this.Visible==false)
            {
                this.Visible = true;
                this.ShowInTaskbar = true;  //显示在系统任务栏 
                this.WindowState = FormWindowState.Normal;  //还原窗体 
                this.mgrNotifyIcon.Visible = true;  //托盘图标隐藏 
            } 
        }

        private void ToolStripMenuItem91_Click(object sender, EventArgs e)
        {
            this.Dispose();
            Application.Exit();
            //this.Close();
        }

        //窗体启动检测
        private void frameMain_Load(object sender, EventArgs e)
        {

            //检查注册表
            checkRegister();
            if (appExist)
            {
                sc_server = new ServiceController("DQS_Server");
                sc_data = new ServiceController("DQS_MongoDB");

                if (sc_server.Status == ServiceControllerStatus.Stopped)
                {
                    serverStatusChange(false);
                }
                else if (sc_server.Status == ServiceControllerStatus.Running)
                {
                    serverStatusChange(true);
                }
                if (sc_data.Status == ServiceControllerStatus.Stopped)
                {
                    dataStatusChange(false);
                }
                else if (sc_data.Status == ServiceControllerStatus.Running)
                {
                    dataStatusChange(true);
                }

                if (localDatabase.Equals("0"))
                {
                    groupBox2.Enabled = false;
                    toolStripMenuItem21.Visible = false;
                    toolStripMenuItem22.Visible = false;
                    toolStripSeparator2.Visible = false;
                }
            }
            else
            {
                groupBox1.Enabled = false; 
                groupBox2.Enabled = false;
                toolStripMenuItem11.Visible = false;
                toolStripMenuItem12.Visible = false;
                toolStripSeparator1.Visible = false; 
                toolStripMenuItem21.Visible = false;
                toolStripMenuItem22.Visible = false;
                toolStripSeparator2.Visible = false;
            }
            

        }


        private void serverStatusChange(bool active)
        {
            if (active == true)
            {
                //更改显示状态
                label1.Text = "running";
                button11.Enabled = false;
                button12.Enabled = true;
                toolStripMenuItem11.Enabled = false;
                toolStripMenuItem12.Enabled = true;
            }
            else
            {
                label1.Text = "stop";
                button11.Enabled = true;
                button12.Enabled = false;
                toolStripMenuItem11.Enabled = true;
                toolStripMenuItem12.Enabled = false;
            }
        }

        private void dataStatusChange(bool active)
        {
            if (active == true)
            {
                //更改显示状态
                label2.Text = "running";
                button21.Enabled = false;
                button22.Enabled = true;
                toolStripMenuItem21.Enabled = false;
                toolStripMenuItem22.Enabled = true;
            }
            else
            {
                label2.Text = "stop";
                button21.Enabled = true;
                button22.Enabled = false;
                toolStripMenuItem21.Enabled = true;
                toolStripMenuItem22.Enabled = false;
            }
        }

        private void serverServiceStart()
        {
            try
            {
                sc_server.Start();
            }
            catch (Exception ex)
            {
                MessageBox.Show("服务启动失败:" + ex.Message);
            }

            serverStatusChange(true);
        }

        private void serverServiceStop()
        {
            try
            {
              sc_server.Stop();
                    //更改显示状态
            }
            catch (Exception ex)
            {
                MessageBox.Show("服务停止失败:" + ex.Message);
            }
            serverStatusChange(false);
        }

        private void dataServiceStart()
        {
            try
            {
                sc_data.Start();
                //更改显示状态
            }
            catch (Exception ex)
            {
                MessageBox.Show("服务启动失败:" + ex.Message);
            }
            dataStatusChange(true);
        }
        private void dataServiceStop()
        {
            try
            {
                  sc_data.Stop();
            }
            catch (Exception ex)
            {
                MessageBox.Show("服务停止失败:" + ex.Message);
            }
            dataStatusChange(false);
        }

        private void button11_Click(object sender, EventArgs e)
        {
            serverServiceStart();
        }

        private void button12_Click(object sender, EventArgs e)
        {
            serverServiceStop();
        }

        private void button21_Click(object sender, EventArgs e)
        {
            dataServiceStart();
        }

        private void button22_Click(object sender, EventArgs e)
        {
            dataServiceStop();
        }

        private void toolStripMenuItem11_Click(object sender, EventArgs e)
        {
            serverServiceStart();
        }

        private void toolStripMenuItem12_Click(object sender, EventArgs e)
        {
            serverServiceStop();
        }

        private void toolStripMenuItem21_Click(object sender, EventArgs e)
        {
            dataServiceStart();
        }

        private void toolStripMenuItem22_Click(object sender, EventArgs e)
        {
            dataServiceStop();
        }

        //检查注册表注册值
        private void checkRegister()
        {
            try
            {
                RegistryKey hkml = Registry.LocalMachine;
                RegistryKey keypath = hkml.OpenSubKey("software\\DqsServer");
                if (keypath != null)
                {
                    appExist = true;
                    localDatabase = keypath.GetValue("database").ToString();
                    installDir = keypath.GetValue("installDir").ToString();
                    keypath.Close();

                }
                else
                {
                    appExist = false;
                }
                
                hkml.Close();
            }
            catch (Exception e)
            {
                localDatabase = "0";
                installDir="C:\\Program Files\\dqs";
            }
            
        }

        private void button13_Click(object sender, EventArgs e)
        {
            Process.Start("explorer.exe", installDir+"\\server\\logs");
        }

        private void button23_Click(object sender, EventArgs e)
        {
            Process.Start("explorer.exe", installDir + "\\data\\logs");
        }

        // 当收到第二个进程的通知时，显示窗体  
        void OnProgramStarted(object state, bool timeout)
        {
            this.mgrNotifyIcon.ShowBalloonTip(1000, "已经运行", "管理器已经在运行", ToolTipIcon.Info);
            this.Show();
            //this.Visible = true;
            //this.ShowInTaskbar = true;  //显示在系统任务栏 
            //this.WindowState = FormWindowState.Normal;  //还原窗体 
            //this.mgrNotifyIcon.Visible = true;  //托盘图标隐藏
            //this.WindowState = FormWindowState.Normal; //注意：一定要在窗体显示后，再对属性进行设置  
        }  
       
    }

}
