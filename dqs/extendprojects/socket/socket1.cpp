#define DLL_API extern "C" _declspec(dllexport)

//#include "stdafx.h"
#include <String.h>
#include <afx.h>
#include "Socket1.h"

//整形数据转化为字符
char MBINT2CHAR(unsigned char ucchar) {
    if((ucchar<=0x09) && (ucchar>=0x00)) {
        return ( '0' + ucchar );
    } else if((ucchar>=0x0A) && (ucchar<=0x0F)) {
        return (ucchar-0x0A+'a');
    } else {
		return '0';
    }
}

//字符转化为整形数据
int MBCHAR2INT(unsigned char ucCharacter) {
    if((ucCharacter>='0') && (ucCharacter<='9')) {
        return (ucCharacter-'0');
    } else if((ucCharacter>='a') && (ucCharacter<='f')) {
        return (ucCharacter-'a'+0x0a);
    } else {
        return 0xFF;
    }
}

//浮点形数据转化为字符串
char *MBFLOAT2CHAR(float sendValue) {
	//协议中定义的float变量基本没有超过999，精度超过小数点后5位的
	//所以我定义了三个整数位，五个小数位
	//以后这个要根据实际应用做相应修改
	int i,n;
	char temp[9];
	//修改为static,静态内存
	static char sendStr[9];
	
	sendValue = sendValue * 100000;
	ltoa(sendValue,temp,10);
	
	
	n = strlen(temp);
	if(n < 8)
	{
		
		for(i=0;i<8-n;i++) {
			sendStr[i] = '0';
		}

		sendStr[i]='\0';
		strcat(sendStr,temp);
	}
	else
	{
		strcpy(sendStr,temp);
	}

	
	/*
	//取整数部分
	temp[0]=(int)sendValue/100;
	temp[1]=(int)(sendValue/10)%10;
	temp[2]=(int)sendValue%10;

	//取小数部分
	//sendValue=sendValue-(int)sendValue;
	temp[3]=(int)(sendValue*10)%10;
	temp[4]=(int)(sendValue*100)%10;
	temp[5]=(int)(sendValue*1000)%10;
	temp[6]=(int)(sendValue*10000)%10;
	temp[7]=(int)(sendValue*100000)%10;
	*/

	/*
	for(i=0;i<8;i++) {
		sendStr[i]='0'+(unsigned char)temp[i];
	}
	*/


	//sendStr[8]='\0';

	return sendStr;
}

//字符串转化为浮点形数据
float MBCHAR2FLOAT(char *recvString) {
	float temp[8];
	int i;
	float recvValue;

	if(strlen(recvString)!=8) {
		return '\0';
	}

	for(i=0;i<8;i++) {
		temp[i]=(float)MBCHAR2INT((unsigned char)(recvString[i]));
	}
	recvValue=temp[0]*100+temp[1]*10+temp[2]+
		temp[3]/10+temp[4]/100+temp[5]/1000+temp[6]/10000+temp[7]/100000;
	return recvValue;
}

unsigned short CRC(unsigned char *puchMsg, unsigned short usDataLen) {
	unsigned char uchCRCHi = 0xFF ; /* 初始化高字节*/
	unsigned char uchCRCLo = 0xFF ; /* 初始化低字节*/
	unsigned uIndex ; /*CRC循环中的索引*/
	while (usDataLen--) {/* 传输消息缓冲区 */
		uIndex = uchCRCHi ^ (*(puchMsg++)) ; /*计算CRC */
		uchCRCHi = uchCRCLo ^ auchCRCHi[uIndex] ;
		uchCRCLo = auchCRCLo[uIndex] ;
	}
	return (uchCRCHi << 8 | uchCRCLo) ;
}

void sendStr(char* b,char *ret) {
	short high8,low8;
	char hhigh4,hlow4,lhigh4,llow4;
	unsigned short code;

	code = CRC((unsigned char *)b,(unsigned short)strlen(b));
	high8=code/(unsigned short)256;
	low8=code%(unsigned short)256;
	hhigh4=(char)high8/(short)16;
	hlow4=(char)high8%(short)16;
	lhigh4=(char)low8/(short)16;
	llow4=(char)low8%(short)16;

	ret[0]=MBINT2CHAR((unsigned char)hhigh4);
	ret[1]=MBINT2CHAR((unsigned char)hlow4);
	ret[2]=MBINT2CHAR((unsigned char)lhigh4);
	ret[3]=MBINT2CHAR((unsigned char)llow4);	
	ret[4] ='\0';
}

//================将明码命令编译成符合协议的命令===========================
//====生成读取参数指令====
//传入参数：
//strParam："设备编号+指令代码"，如"SI01234567g"
//retParam：返回的读取参数指令字符串
bool GenerateReadParam(char *strParam,char *retParam) {
	char str1[5];

	strcpy(retParam, strParam);
	strcat(retParam,"0000");
	sendStr(retParam,str1);
	strcat(retParam,str1);

	return true;
}

//====生成设置参数命令====
//传入参数：
//strParam："设备编号+指令代码"，如"SI01234567s"
//setStruct：设置参数结构体
//retParam：返回的设置参数指令字符串
bool GenerateSetParam(char *strParam, retData &setStruct, char *retParam) {
	int i;
	char temp[11];
 	//char str1[5];
	//fjh 20120907

	 	
 	//添加数据长度字符串
 	strcpy(retParam,strParam);
 	strcat(retParam,"005F");
 
 	//添加具体数据字符串
 	//基本参数
	if(strlen(setStruct._chuanGanQiNumber) == 10) {
 		strcat(retParam,setStruct._chuanGanQiNumber);//传感器编号
	} else if(strlen(setStruct._chuanGanQiNumber) < 10) {
		strcat(retParam,setStruct._chuanGanQiNumber);
		for(i=strlen(setStruct._chuanGanQiNumber);i<10;i++) {
			strcat(retParam," ");
		}
	} else {
		for(i=0;i<10;i++) {
			temp[i] = setStruct._chuanGanQiNumber[i];
		}
		temp[10]='\0';
		strcat(retParam,temp);
	}
	if(strlen(setStruct._zhanDian) == 10) {
 		strcat(retParam,setStruct._zhanDian);//站点名称
	} else if(strlen(setStruct._zhanDian) < 10) {
		strcat(retParam,setStruct._zhanDian);
		for(i=strlen(setStruct._zhanDian);i<10;i++) {
			strcat(retParam," ");
		}
	} else {
		for(i=0;i<10;i++) {
			temp[i] = setStruct._zhanDian[i];
		}
		temp[10]='\0';
		strcat(retParam,temp);
	}
 	
	
	strcat(retParam,MBFLOAT2CHAR(setStruct._jingDu));//经度
	strcat(retParam,MBFLOAT2CHAR(setStruct._weiDu));//纬度

 	retParam[51]=MBINT2CHAR((unsigned char)setStruct._changDi);//场地类型
 	retParam[52]=MBINT2CHAR((unsigned char)setStruct._guanCe);//观测对象
 	retParam[53]=MBINT2CHAR((unsigned char)setStruct._jiaSuDuJi);//加速度计型号
 	retParam[54]=MBINT2CHAR((unsigned char)setStruct._anZhuangFangXiang);//安装方向
 	retParam[55]=MBINT2CHAR((unsigned char)setStruct._liangCheng);//量程选择
 	retParam[56]='\0';
	strcat(retParam,MBFLOAT2CHAR(setStruct._caiYang));//采样周期
	if(strlen(setStruct._xingZheng) == 6) {
		strcat(retParam,setStruct._xingZheng);//行政区域编码
	} else if(strlen(setStruct._xingZheng) < 6) {
		strcat(retParam,setStruct._xingZheng);
		for(i=strlen(setStruct._xingZheng);i<6;i++) {
			strcat(retParam," ");
		}
	} else {
		for(i=0;i<6;i++) {
			temp[i] = setStruct._xingZheng[i];
		}
		temp[6]='\0';
		strcat(retParam,temp);
	}
	if(strlen(setStruct._yuLiu1) == 8) {
		strcat(retParam,setStruct._yuLiu1);//预留1
	} else if(strlen(setStruct._yuLiu1) < 8) {
		strcat(retParam,setStruct._yuLiu1);
		for(i=strlen(setStruct._yuLiu1);i<8;i++) {
			strcat(retParam," ");
		}
	} else {
		for(i=0;i<8;i++) {
			temp[i] = setStruct._yuLiu1[i];
		}
		temp[8]='\0';
		strcat(retParam,temp);
	}
	if(strlen(setStruct._yuLiu2) == 8) {
		strcat(retParam,setStruct._yuLiu2);//预留2
	} else if(strlen(setStruct._yuLiu2) < 8) {
		strcat(retParam,setStruct._yuLiu2);
		for(i=strlen(setStruct._yuLiu2);i<8;i++) {
			strcat(retParam," ");
		}
	} else {
		for(i=0;i<8;i++) {
			temp[i] = setStruct._yuLiu2[i];
		}
		temp[8]='\0';
		strcat(retParam,temp);
	}
 
 	//触发参数
 	retParam[86]=MBINT2CHAR((unsigned char)setStruct._pgaChufa);//PGA触发
 	retParam[87]='\0';
 	strcat(retParam,MBFLOAT2CHAR(setStruct._pgaYuzhi));//PGA阈值
 	retParam[95]=MBINT2CHAR((unsigned char)setStruct._siChufa);//SI触发
 	retParam[96]='\0';
 	strcat(retParam,MBFLOAT2CHAR(setStruct._siYuZhi));//SI阈值
 	retParam[104]=MBINT2CHAR((unsigned char)setStruct._zuHeChufa);//组合触发
 	retParam[105]=MBINT2CHAR((unsigned char)setStruct._yuLiuChufa);//预留触发
 	retParam[106]='\0';
 	strcat(retParam,MBFLOAT2CHAR(setStruct._yuLiuYuzhi));//预留阈值
 	
 	//报警参数
 	retParam[114]=MBINT2CHAR((unsigned char)setStruct._alarmPGA);//PGA报警
 	retParam[115]='\0';
 	strcat(retParam,MBFLOAT2CHAR(setStruct._alarmPGAYuzhi));//PGA阈值
 	retParam[123]=MBINT2CHAR((unsigned char)setStruct._alarmSI);//SI报警
 	retParam[124]='\0';
 	strcat(retParam,MBFLOAT2CHAR(setStruct._alarmSIYuzhi));//SI阈值
 	retParam[132]=MBINT2CHAR((unsigned char)setStruct._alarmGroup);//组合报警
 	retParam[133]=MBINT2CHAR((unsigned char)setStruct._alarmYuliu);//预留报警
 	retParam[134]='\0';
 	strcat(retParam,MBFLOAT2CHAR(setStruct._alarmYUliuYuzhi));//预留阈值
 	
 	//输出参数
 	retParam[142]=MBINT2CHAR((unsigned char)setStruct._dAOutput1);//D/A输出1
 	retParam[143]=MBINT2CHAR((unsigned char)setStruct._dAOutput2);//D/A输出2
 	retParam[144]=MBINT2CHAR((unsigned char)setStruct._IO1);//I/O输出1
 	retParam[145]=MBINT2CHAR((unsigned char)setStruct._IO2);//I/O输出2
 	retParam[146]='\0';
 	
 	//添加crc校验
 //	sendStr(strParam,str1);
 //	strcat(retParam,str1);
	
	return true;
}

//====生成突发数据字符串====
//传入参数：
//strParam："设备编号+指令代码"，如"SI01234567a"
//setStruct：设置参数结构体
//retParam：返回的设置参数指令字符串
bool GenerateFlashParam(char *strParam, FlashData &flashStruct, char *retParam) {
// 	char str1[5];
//	char str2[11];
	int i;
 	
 	//添加数据长度字符串
 	strcpy(retParam, strParam);
 	strcat(retParam, "003B");//15字节
	//strcpy(retParam,"SI30001001s005F");
 
 	//添加具体数据字符串
 	//基本参数
	if(strlen(flashStruct._jilubianhao) < 10) {
		for(i=strlen(flashStruct._jilubianhao); i<10; i++){
			flashStruct._jilubianhao[i]=' ';
		}
	}
	strcat(retParam,flashStruct._jilubianhao);//记录编号
	retParam[25]='\0';//25

	if(strlen(flashStruct._chuanGanQiNumber) < 10) {
		for(i=strlen(flashStruct._chuanGanQiNumber);i<10;i++){
			flashStruct._chuanGanQiNumber[i]=' ';
		}
	}
 	strcat(retParam,flashStruct._chuanGanQiNumber);//传感器编号
	retParam[25]='\0';//35

	strcat(retParam,MBFLOAT2CHAR(flashStruct._jingDu));//经度43
	strcat(retParam,MBFLOAT2CHAR(flashStruct._weiDu));//纬度51
 	retParam[51]=MBINT2CHAR((unsigned char)flashStruct._changDi);//场地类型52
 	retParam[52]=MBINT2CHAR((unsigned char)flashStruct._guanCe);//观测对象53
 	retParam[53]=MBINT2CHAR((unsigned char)flashStruct._anZhuangFangXiang);//安装方向54
 	retParam[54]='\0';
	if(strlen(flashStruct._xingZheng) < 6) {
		for(i=strlen(flashStruct._xingZheng);i<6;i++) {
			flashStruct._xingZheng[i] = ' ';//不足6位，用空格补齐
		}
	}
 	strcat(retParam,flashStruct._xingZheng);//行政区域编码60
	retParam[60]='\0';

	if(strlen(flashStruct._chushishike) < 6) {
		for(i=strlen(flashStruct._chushishike);i<6;i++) {
			flashStruct._chushishike[i] = ' ';//不足6位，用空格补齐
		}
	}
//	strcat(retParam,flashStruct._chushishike);
//	retParam[66]='\0';
	for(i=0; i<6; i++) {
		retParam[60+i*2] = MBINT2CHAR((unsigned char) (flashStruct._chushishike[i] / 16));
		retParam[60+i*2+1] = MBINT2CHAR((unsigned char) (flashStruct._chushishike[i] % 16));
	}
	retParam[72] = '\0';//初始时刻72

	strcat(retParam,MBFLOAT2CHAR(flashStruct._caiYang));//采样周期80
 
 	strcat(retParam,MBFLOAT2CHAR(flashStruct._pgazhi));//PGA阈值88
 	strcat(retParam,MBFLOAT2CHAR(flashStruct._siZhi));//SI阈值96
	for(i=0;i<7;i++) {//记录长度103
		retParam[96+i] = MBINT2CHAR((unsigned char)((long)flashStruct._jiluchangdu%10));
	}
	retParam[103] = '\0';
 	
 	//添加crc校验
 //	sendStr(strParam,str1);
 //	strcat(retParam,str1);
	
	return true;
}

//========================解析读取到的设置参数=============================
//recvParam：socket直接返回的字符串
//retParam：返回的结构体
bool parseReadSetParam(char *recvParam,retData &set) {
	//临时变量
//	int i=0;
	char temp[1024];

//	if(strlen(recvParam)!=146)//判断字符长度
//	{
//		return false;
//	}
	char temp1[1024];

	//验证CRC
//	for(i=0;i<strlen(recvParam)-4;i++)
//	{
//		temp1[i]=recvParam[i];
//	}
//	char str2[5];
//	sendStr(temp1,str2);

//	strcat(temp1,str2);
//	if(strcmp(recvParam,temp1))
//	{
//		return false;
//	}

	//验证完毕,对结构体赋值
	//获得读取回来的设置参数字符串
//	char readParam[131];
//	for (int i=0;i<STRUCT2STRING;i++)
//	{
		//接收到的字符串格式为：10字节设备编码、1字节命令参数、4字节参数长度、
		//131字节具体参数、4字节crc校验
//		readParam[i]=recvParam[i+15];
//	}
	
	int i;
	//基本参数赋值
	for(i=0;i<10;i++) {//传感器编号
		temp1[i]=recvParam[i+15];
	}
	temp1[10]='\0';
	strcpy(set._chuanGanQiNumber,temp1);
	for(i=0;i<10;i++) {//站点名称
		temp1[i]=recvParam[i+25];
	}
	temp1[10]='\0';
	strcpy(set._zhanDian,temp1);
	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+35];
	}
	temp[8]='\0';
	set._jingDu=MBCHAR2FLOAT(temp);//经度
	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+43];
	}
	temp[8]='\0';
	set._weiDu=MBCHAR2FLOAT(temp);//纬度
	set._changDi=MBCHAR2INT((unsigned char)recvParam[51]);//场地类型
	set._guanCe=MBCHAR2INT((unsigned char)recvParam[52]);//观测对象
	set._jiaSuDuJi=MBCHAR2INT((unsigned char)recvParam[53]);//加速度计型号
	set._anZhuangFangXiang=MBCHAR2INT((unsigned char)recvParam[54]);//安装方向
	set._liangCheng=MBCHAR2INT((unsigned char)recvParam[55]);//量程选择
	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+56];
	}
	temp[8]='\0';
	set._caiYang=MBCHAR2FLOAT(temp);//采样周期
	//行政区域编码
	for(i=0;i<6;i++) {
		temp1[i]=recvParam[64+i];
	}
	temp1[6]='\0';
	strcpy(set._xingZheng,temp1);
	//预留1
	for(i=0;i<8;i++) {
		temp1[i]=recvParam[70+i];
	}
	temp1[8]='\0';
	strcpy(set._yuLiu1,temp1);
	//预留2
	for(i=0;i<8;i++) {
		temp1[i]=recvParam[78+i];
	}
	temp1[8]='\0';
	strcpy(set._yuLiu2,temp1);

	//触发参数赋值
	set._pgaChufa=MBCHAR2INT((unsigned char)recvParam[86]);//PGA触发
	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+87];
	}
	temp[8]='\0';
	set._pgaYuzhi=MBCHAR2FLOAT(temp);//PGA阈值
	set._siChufa=MBCHAR2INT((unsigned char)recvParam[95]);//SI触发
	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+96];
	}
	temp[8]='\0';
	set._siYuZhi=MBCHAR2FLOAT(temp);//SI阈值
	set._zuHeChufa=MBCHAR2INT((unsigned char)recvParam[104]);//组合触发
	set._yuLiuChufa=MBCHAR2INT((unsigned char)recvParam[105]);//预留触发
	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+106];
	}
	temp[8]='\0';
	set._yuLiuYuzhi=MBCHAR2FLOAT(temp);//预留阈值

	//报警参数赋值
	set._alarmPGA=MBCHAR2INT((unsigned char)recvParam[114]);//PGA报警
	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+115];
	}
	temp[8]='\0';
	set._alarmPGAYuzhi=MBCHAR2FLOAT(temp);//PGA阈值
	set._alarmSI=MBCHAR2INT((unsigned char)recvParam[123]);//SI报警
	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+124];
	}
	temp[8]='\0';
	set._alarmSIYuzhi=MBCHAR2FLOAT(temp);//SI阈值
	set._alarmGroup=MBCHAR2INT((unsigned char)recvParam[132]);//组合报警
	set._alarmYuliu=MBCHAR2INT((unsigned char)recvParam[133]);//预留报警
	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+134];
	}
	temp[8]='\0';
	set._alarmYUliuYuzhi=MBCHAR2FLOAT(temp);//预留阈值

	//报警参数赋值
	set._dAOutput1=MBCHAR2INT((unsigned char)recvParam[142]);//D/A输出1
	set._dAOutput2=MBCHAR2INT((unsigned char)recvParam[143]);//D/A输出2
	set._IO1=MBCHAR2INT((unsigned char)recvParam[144]);//I/O输出1
	set._IO2=MBCHAR2INT((unsigned char)recvParam[145]);//I/O输出2

	return true;
}

//========================解析读取到的突发参数=============================
//recvParam：socket直接返回的字符串
//retParam：返回的结构体
bool parseReadFlashParam(char *recvParam,FlashData &set) {
	//临时变量
	int i=0;
	char temp[1024];
	char temp1[1024];
	int length[7];

	//验证CRC
//	for(i=0;i<strlen(recvParam)-4;i++)
//	{
//		temp1[i]=recvParam[i];
//	}
//	char str2[5];
//	sendStr(temp1,str2);

//	strcat(temp1,str2);
//	if(strcmp(recvParam,temp1))
//	{
//		return false;
//	}

	//验证完毕,对结构体赋值
	//获得读取回来的设置参数字符串
//	char readParam[131];
//	for (int i=0;i<STRUCT2STRING;i++)
//	{
		//接收到的字符串格式为：10字节设备编码、1字节命令参数、4字节参数长度、
		//131字节具体参数、4字节crc校验
//		readParam[i]=recvParam[i+15];
//	}
	
	//基本参数赋值
	for(i=0;i<10;i++) {//记录编号
		temp1[i]=recvParam[i+15];
	}
	temp1[10]='\0';
	strcpy(set._jilubianhao,temp1);

	for(i=0;i<10;i++) {//传感器编号
		temp1[i]=recvParam[i+25];
	}
	temp1[10]='\0';
	strcpy(set._chuanGanQiNumber,temp1);

	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+35];
	}
	temp[8]='\0';
	set._jingDu=MBCHAR2FLOAT(temp);//经度
	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+43];
	}
	temp[8]='\0';
	set._weiDu=MBCHAR2FLOAT(temp);//纬度
	set._changDi=MBCHAR2INT((unsigned char)recvParam[51]);//场地类型
	set._guanCe=MBCHAR2INT((unsigned char)recvParam[52]);//观测对象
	set._anZhuangFangXiang=MBCHAR2INT((unsigned char)recvParam[53]);//安装方向

	//行政区域编码
	for(i=0;i<6;i++) {
		temp1[i]=recvParam[54+i];
	}
	temp1[6]='\0';
	strcpy(set._xingZheng,temp1);

	//初始时刻
	for(i=0;i<6;i++) {
		temp1[i]=(char)((MBCHAR2INT(recvParam[60+i*2]) << 4) + MBCHAR2INT(recvParam[60+i*2+1]));
	}
	temp1[6]='\0';
	strcpy(set._chushishike,temp1);

	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+72];
	}
	temp[8]='\0';
	set._caiYang=MBCHAR2FLOAT(temp);//采样周期
	
	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+80];
	}
	temp[8]='\0';
	set._pgazhi=MBCHAR2FLOAT(temp);//PGA值

	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+88];
	}
	temp[8]='\0';
	set._siZhi=MBCHAR2FLOAT(temp);//SI值

	//纪录长度
	for(i=0;i<7;i++) {
		length[i] = MBCHAR2INT((unsigned char)recvParam[i+96]);
	}
	set._jiluchangdu = length[6]*1000000+length[5]*100000+length[4]*10000+length[3]*1000+length[2]*100+length[1]*10+length[0];

	return true;
}

//解析删除Flash中存储数据指令
//传入参数：
//recvParam：socket直接返回的字符串
bool ParseDelParam(char *recvParam) {
	int i;
	char result;
	char *temp1="";
	char str2[5];

	//验证CRC
	for(i=0;i<(strlen(recvParam)-4);i++) {
		temp1[i]=recvParam[i];
	}
	
	sendStr(temp1,str2);

	strcat(temp1,str2);
	if(strcmp(recvParam,temp1)) {
		return false;
	}

	//验证完毕，返回结果
	result=recvParam[15];
	
	//0x01时表示设置成功，0x00时表示设置失败
	//0x01在socket传输时被转化为字符'1'
	//0x00在socket传输时被转化为字符'0'
	if(!(result-'1')) {
		return true;
	} else {
		return false;
	}
}


//解析设置参数指令的返回值
//传入参数：
//recvParam：socket直接返回的字符串
bool ParseSetParam(char *recvParam) {
	char result;
	
	/*
	char *temp1="";
	char str2[5];
	int i;

	//验证CRC
	for(i=0;i<(strlen(recvParam)-4);i++) {
		temp1[i]=recvParam[i];
	}
	temp1[i]='\0';
	
	//str2=sendStr(temp1);
	sendStr(temp1,str2);

	strcat(temp1,str2);
	if(strcmp(recvParam,temp1)) {
		return false;
	}
	*/
	//验证完毕，返回结果
	result=recvParam[15];

	//协议规定：0x01时表示设置成功，0x00时表示设置失败
	//0x01在socket传输时被转化为字符'1'
	//0x00在socket传输时被转化为字符'0'
	if(!(result-'1')) {
		//当result为1时
		return true;
	} else {
		return false;
	}
}


//读取存储在Flash中传感器的数据指令
//分包发送，包的格式为：
//========================================================================================
//=帧头（固定：0x00H）、地址（设备编号）、功能码（固定：'r'）、数据大小（根据计算得到）、=
//=数据字节数（具体的要传输的数据）、CRC校验位、帧尾（固定：0x7FH）                      =
//========================================================================================
//服务器端接解析函数
//recvFlashData: 传入参数，从设备接收到的字符串（具体内容见上述“包的格式”）
//ret: 返回参数，一个浮点型数组
bool parseRecvFlashData(char *recvFlashData, float *ret) {
	int i,temp;
	int recvDataLen=0;

	//判断帧头是否'00'
	if(recvFlashData[0]!='0' || recvFlashData[1]!='0') {
		return false;
	}
	//判断功能码是否'r'
	if(recvFlashData[12]!='r') {
		return false;
	}

	//检验数据的大小是否正确
	for(i=13;i<21;i++) {
		temp=MBCHAR2INT((unsigned char)recvFlashData[i]);

		recvDataLen+=(temp<<((i-13)*4));
	}
	if((recvDataLen=+27) != strlen(recvFlashData)) {
		return false;
	}

	//CRC校验
	

	//开始解析
	int flag;
	float f;

	for(temp=0;temp<recvDataLen/9;temp++) {
		f=0;
		flag=MBCHAR2INT((unsigned char)recvFlashData[21+temp*9]);

		for(i=0;i<8;i++) {
			f+=((MBCHAR2INT((unsigned char)recvFlashData[21+temp*9+i+1]))<<4*i);
		}
		if(flag==1) {
			ret[temp]=0-f/1000000;
		} else {
			ret[temp]=f/1000000;
		}
	}

	return true;
}

//生成Flash数据
//deviceNum：设备编号
//readData[]：从烈度计收到的数据
//frameNum：帧数
//retString：生成的字符串
bool generateFlashData(char *deviceNum, short readData[480], short frameNum, char *retString) {
	char temp[1024];
	int i = 0;
	char str2[5];

	strcpy(temp, deviceNum);
	strcat(temp, "r01E0");

	for(i=0; i<480; i++) {
		//这个地方的长度写死了，以后根据不同的采样频率可能还会修改
		temp[15+i*2] = MBINT2CHAR((unsigned char)(readData[i]/16));
		temp[15+i*2+1] = MBINT2CHAR((unsigned char)(readData[i]%16));
	}
	
	temp[975] = '\0';
	
	//crc校验
	sendStr(temp,str2);
	strcat(temp,str2);

	temp[979] = MBINT2CHAR((unsigned char)(frameNum/16));
	temp[980] = MBINT2CHAR((unsigned char)(frameNum%16));
	temp[981] = '\0';

	return true;
}

//解析Flash数据
//receiveStr：收到的字符串
//readData[]：解析出的数据
//deviceID：设备编号
//返回值：帧数
short parseFlashData(char *receiveStr, short readData[240], char *deviceID) {
	int i;
	int frameNum;
	char temp[11];
	short tempShort;

	for(i=0;i<10;i++) {
		temp[i] = receiveStr[i];
	}
	temp[10] = '\0';
	//strcmp(deviceID, temp);

	for(i=0; i<240; i++) {
		tempShort = ((short)MBCHAR2INT((unsigned char)receiveStr[15+i*4]))*16 + (short)MBCHAR2INT((unsigned char)receiveStr[15+i*4+1]) +
			256*((short)MBCHAR2INT((unsigned char)receiveStr[15+i*4+2])*16 + (short)MBCHAR2INT((unsigned char)receiveStr[15+i*4+3]));
		if((tempShort & 0x8000) == 0x8000) {
			readData[i] = tempShort - 65536;
		} else {
			readData[i] = tempShort;
		}
	}

	frameNum = (short)MBCHAR2INT((unsigned char)receiveStr[979]) * 16 + (short)MBCHAR2INT((unsigned char)receiveStr[980]);

	return frameNum;
}