#define DLL_API extern "C" _declspec(dllexport)

//#include "stdafx.h"
#include <String.h>
#include <afx.h>
#include "Socket1.h"

//��������ת��Ϊ�ַ�
char MBINT2CHAR(unsigned char ucchar) {
    if((ucchar<=0x09) && (ucchar>=0x00)) {
        return ( '0' + ucchar );
    } else if((ucchar>=0x0A) && (ucchar<=0x0F)) {
        return (ucchar-0x0A+'a');
    } else {
		return '0';
    }
}

//�ַ�ת��Ϊ��������
int MBCHAR2INT(unsigned char ucCharacter) {
    if((ucCharacter>='0') && (ucCharacter<='9')) {
        return (ucCharacter-'0');
    } else if((ucCharacter>='a') && (ucCharacter<='f')) {
        return (ucCharacter-'a'+0x0a);
    } else {
        return 0xFF;
    }
}

//����������ת��Ϊ�ַ���
char *MBFLOAT2CHAR(float sendValue) {
	//Э���ж����float��������û�г���999�����ȳ���С�����5λ��
	//�����Ҷ�������������λ�����С��λ
	//�Ժ����Ҫ����ʵ��Ӧ������Ӧ�޸�
	int i,n;
	char temp[9];
	//�޸�Ϊstatic,��̬�ڴ�
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
	//ȡ��������
	temp[0]=(int)sendValue/100;
	temp[1]=(int)(sendValue/10)%10;
	temp[2]=(int)sendValue%10;

	//ȡС������
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

//�ַ���ת��Ϊ����������
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
	unsigned char uchCRCHi = 0xFF ; /* ��ʼ�����ֽ�*/
	unsigned char uchCRCLo = 0xFF ; /* ��ʼ�����ֽ�*/
	unsigned uIndex ; /*CRCѭ���е�����*/
	while (usDataLen--) {/* ������Ϣ������ */
		uIndex = uchCRCHi ^ (*(puchMsg++)) ; /*����CRC */
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

//================�������������ɷ���Э�������===========================
//====���ɶ�ȡ����ָ��====
//���������
//strParam��"�豸���+ָ�����"����"SI01234567g"
//retParam�����صĶ�ȡ����ָ���ַ���
bool GenerateReadParam(char *strParam,char *retParam) {
	char str1[5];

	strcpy(retParam, strParam);
	strcat(retParam,"0000");
	sendStr(retParam,str1);
	strcat(retParam,str1);

	return true;
}

//====�������ò�������====
//���������
//strParam��"�豸���+ָ�����"����"SI01234567s"
//setStruct�����ò����ṹ��
//retParam�����ص����ò���ָ���ַ���
bool GenerateSetParam(char *strParam, retData &setStruct, char *retParam) {
	int i;
	char temp[11];
 	//char str1[5];
	//fjh 20120907

	 	
 	//������ݳ����ַ���
 	strcpy(retParam,strParam);
 	strcat(retParam,"005F");
 
 	//��Ӿ��������ַ���
 	//��������
	if(strlen(setStruct._chuanGanQiNumber) == 10) {
 		strcat(retParam,setStruct._chuanGanQiNumber);//���������
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
 		strcat(retParam,setStruct._zhanDian);//վ������
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
 	
	
	strcat(retParam,MBFLOAT2CHAR(setStruct._jingDu));//����
	strcat(retParam,MBFLOAT2CHAR(setStruct._weiDu));//γ��

 	retParam[51]=MBINT2CHAR((unsigned char)setStruct._changDi);//��������
 	retParam[52]=MBINT2CHAR((unsigned char)setStruct._guanCe);//�۲����
 	retParam[53]=MBINT2CHAR((unsigned char)setStruct._jiaSuDuJi);//���ٶȼ��ͺ�
 	retParam[54]=MBINT2CHAR((unsigned char)setStruct._anZhuangFangXiang);//��װ����
 	retParam[55]=MBINT2CHAR((unsigned char)setStruct._liangCheng);//����ѡ��
 	retParam[56]='\0';
	strcat(retParam,MBFLOAT2CHAR(setStruct._caiYang));//��������
	if(strlen(setStruct._xingZheng) == 6) {
		strcat(retParam,setStruct._xingZheng);//�����������
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
		strcat(retParam,setStruct._yuLiu1);//Ԥ��1
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
		strcat(retParam,setStruct._yuLiu2);//Ԥ��2
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
 
 	//��������
 	retParam[86]=MBINT2CHAR((unsigned char)setStruct._pgaChufa);//PGA����
 	retParam[87]='\0';
 	strcat(retParam,MBFLOAT2CHAR(setStruct._pgaYuzhi));//PGA��ֵ
 	retParam[95]=MBINT2CHAR((unsigned char)setStruct._siChufa);//SI����
 	retParam[96]='\0';
 	strcat(retParam,MBFLOAT2CHAR(setStruct._siYuZhi));//SI��ֵ
 	retParam[104]=MBINT2CHAR((unsigned char)setStruct._zuHeChufa);//��ϴ���
 	retParam[105]=MBINT2CHAR((unsigned char)setStruct._yuLiuChufa);//Ԥ������
 	retParam[106]='\0';
 	strcat(retParam,MBFLOAT2CHAR(setStruct._yuLiuYuzhi));//Ԥ����ֵ
 	
 	//��������
 	retParam[114]=MBINT2CHAR((unsigned char)setStruct._alarmPGA);//PGA����
 	retParam[115]='\0';
 	strcat(retParam,MBFLOAT2CHAR(setStruct._alarmPGAYuzhi));//PGA��ֵ
 	retParam[123]=MBINT2CHAR((unsigned char)setStruct._alarmSI);//SI����
 	retParam[124]='\0';
 	strcat(retParam,MBFLOAT2CHAR(setStruct._alarmSIYuzhi));//SI��ֵ
 	retParam[132]=MBINT2CHAR((unsigned char)setStruct._alarmGroup);//��ϱ���
 	retParam[133]=MBINT2CHAR((unsigned char)setStruct._alarmYuliu);//Ԥ������
 	retParam[134]='\0';
 	strcat(retParam,MBFLOAT2CHAR(setStruct._alarmYUliuYuzhi));//Ԥ����ֵ
 	
 	//�������
 	retParam[142]=MBINT2CHAR((unsigned char)setStruct._dAOutput1);//D/A���1
 	retParam[143]=MBINT2CHAR((unsigned char)setStruct._dAOutput2);//D/A���2
 	retParam[144]=MBINT2CHAR((unsigned char)setStruct._IO1);//I/O���1
 	retParam[145]=MBINT2CHAR((unsigned char)setStruct._IO2);//I/O���2
 	retParam[146]='\0';
 	
 	//���crcУ��
 //	sendStr(strParam,str1);
 //	strcat(retParam,str1);
	
	return true;
}

//====����ͻ�������ַ���====
//���������
//strParam��"�豸���+ָ�����"����"SI01234567a"
//setStruct�����ò����ṹ��
//retParam�����ص����ò���ָ���ַ���
bool GenerateFlashParam(char *strParam, FlashData &flashStruct, char *retParam) {
// 	char str1[5];
//	char str2[11];
	int i;
 	
 	//������ݳ����ַ���
 	strcpy(retParam, strParam);
 	strcat(retParam, "003B");//15�ֽ�
	//strcpy(retParam,"SI30001001s005F");
 
 	//��Ӿ��������ַ���
 	//��������
	if(strlen(flashStruct._jilubianhao) < 10) {
		for(i=strlen(flashStruct._jilubianhao); i<10; i++){
			flashStruct._jilubianhao[i]=' ';
		}
	}
	strcat(retParam,flashStruct._jilubianhao);//��¼���
	retParam[25]='\0';//25

	if(strlen(flashStruct._chuanGanQiNumber) < 10) {
		for(i=strlen(flashStruct._chuanGanQiNumber);i<10;i++){
			flashStruct._chuanGanQiNumber[i]=' ';
		}
	}
 	strcat(retParam,flashStruct._chuanGanQiNumber);//���������
	retParam[25]='\0';//35

	strcat(retParam,MBFLOAT2CHAR(flashStruct._jingDu));//����43
	strcat(retParam,MBFLOAT2CHAR(flashStruct._weiDu));//γ��51
 	retParam[51]=MBINT2CHAR((unsigned char)flashStruct._changDi);//��������52
 	retParam[52]=MBINT2CHAR((unsigned char)flashStruct._guanCe);//�۲����53
 	retParam[53]=MBINT2CHAR((unsigned char)flashStruct._anZhuangFangXiang);//��װ����54
 	retParam[54]='\0';
	if(strlen(flashStruct._xingZheng) < 6) {
		for(i=strlen(flashStruct._xingZheng);i<6;i++) {
			flashStruct._xingZheng[i] = ' ';//����6λ���ÿո���
		}
	}
 	strcat(retParam,flashStruct._xingZheng);//�����������60
	retParam[60]='\0';

	if(strlen(flashStruct._chushishike) < 6) {
		for(i=strlen(flashStruct._chushishike);i<6;i++) {
			flashStruct._chushishike[i] = ' ';//����6λ���ÿո���
		}
	}
//	strcat(retParam,flashStruct._chushishike);
//	retParam[66]='\0';
	for(i=0; i<6; i++) {
		retParam[60+i*2] = MBINT2CHAR((unsigned char) (flashStruct._chushishike[i] / 16));
		retParam[60+i*2+1] = MBINT2CHAR((unsigned char) (flashStruct._chushishike[i] % 16));
	}
	retParam[72] = '\0';//��ʼʱ��72

	strcat(retParam,MBFLOAT2CHAR(flashStruct._caiYang));//��������80
 
 	strcat(retParam,MBFLOAT2CHAR(flashStruct._pgazhi));//PGA��ֵ88
 	strcat(retParam,MBFLOAT2CHAR(flashStruct._siZhi));//SI��ֵ96
	for(i=0;i<7;i++) {//��¼����103
		retParam[96+i] = MBINT2CHAR((unsigned char)((long)flashStruct._jiluchangdu%10));
	}
	retParam[103] = '\0';
 	
 	//���crcУ��
 //	sendStr(strParam,str1);
 //	strcat(retParam,str1);
	
	return true;
}

//========================������ȡ�������ò���=============================
//recvParam��socketֱ�ӷ��ص��ַ���
//retParam�����صĽṹ��
bool parseReadSetParam(char *recvParam,retData &set) {
	//��ʱ����
//	int i=0;
	char temp[1024];

//	if(strlen(recvParam)!=146)//�ж��ַ�����
//	{
//		return false;
//	}
	char temp1[1024];

	//��֤CRC
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

	//��֤���,�Խṹ�帳ֵ
	//��ö�ȡ���������ò����ַ���
//	char readParam[131];
//	for (int i=0;i<STRUCT2STRING;i++)
//	{
		//���յ����ַ�����ʽΪ��10�ֽ��豸���롢1�ֽ����������4�ֽڲ������ȡ�
		//131�ֽھ��������4�ֽ�crcУ��
//		readParam[i]=recvParam[i+15];
//	}
	
	int i;
	//����������ֵ
	for(i=0;i<10;i++) {//���������
		temp1[i]=recvParam[i+15];
	}
	temp1[10]='\0';
	strcpy(set._chuanGanQiNumber,temp1);
	for(i=0;i<10;i++) {//վ������
		temp1[i]=recvParam[i+25];
	}
	temp1[10]='\0';
	strcpy(set._zhanDian,temp1);
	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+35];
	}
	temp[8]='\0';
	set._jingDu=MBCHAR2FLOAT(temp);//����
	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+43];
	}
	temp[8]='\0';
	set._weiDu=MBCHAR2FLOAT(temp);//γ��
	set._changDi=MBCHAR2INT((unsigned char)recvParam[51]);//��������
	set._guanCe=MBCHAR2INT((unsigned char)recvParam[52]);//�۲����
	set._jiaSuDuJi=MBCHAR2INT((unsigned char)recvParam[53]);//���ٶȼ��ͺ�
	set._anZhuangFangXiang=MBCHAR2INT((unsigned char)recvParam[54]);//��װ����
	set._liangCheng=MBCHAR2INT((unsigned char)recvParam[55]);//����ѡ��
	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+56];
	}
	temp[8]='\0';
	set._caiYang=MBCHAR2FLOAT(temp);//��������
	//�����������
	for(i=0;i<6;i++) {
		temp1[i]=recvParam[64+i];
	}
	temp1[6]='\0';
	strcpy(set._xingZheng,temp1);
	//Ԥ��1
	for(i=0;i<8;i++) {
		temp1[i]=recvParam[70+i];
	}
	temp1[8]='\0';
	strcpy(set._yuLiu1,temp1);
	//Ԥ��2
	for(i=0;i<8;i++) {
		temp1[i]=recvParam[78+i];
	}
	temp1[8]='\0';
	strcpy(set._yuLiu2,temp1);

	//����������ֵ
	set._pgaChufa=MBCHAR2INT((unsigned char)recvParam[86]);//PGA����
	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+87];
	}
	temp[8]='\0';
	set._pgaYuzhi=MBCHAR2FLOAT(temp);//PGA��ֵ
	set._siChufa=MBCHAR2INT((unsigned char)recvParam[95]);//SI����
	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+96];
	}
	temp[8]='\0';
	set._siYuZhi=MBCHAR2FLOAT(temp);//SI��ֵ
	set._zuHeChufa=MBCHAR2INT((unsigned char)recvParam[104]);//��ϴ���
	set._yuLiuChufa=MBCHAR2INT((unsigned char)recvParam[105]);//Ԥ������
	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+106];
	}
	temp[8]='\0';
	set._yuLiuYuzhi=MBCHAR2FLOAT(temp);//Ԥ����ֵ

	//����������ֵ
	set._alarmPGA=MBCHAR2INT((unsigned char)recvParam[114]);//PGA����
	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+115];
	}
	temp[8]='\0';
	set._alarmPGAYuzhi=MBCHAR2FLOAT(temp);//PGA��ֵ
	set._alarmSI=MBCHAR2INT((unsigned char)recvParam[123]);//SI����
	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+124];
	}
	temp[8]='\0';
	set._alarmSIYuzhi=MBCHAR2FLOAT(temp);//SI��ֵ
	set._alarmGroup=MBCHAR2INT((unsigned char)recvParam[132]);//��ϱ���
	set._alarmYuliu=MBCHAR2INT((unsigned char)recvParam[133]);//Ԥ������
	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+134];
	}
	temp[8]='\0';
	set._alarmYUliuYuzhi=MBCHAR2FLOAT(temp);//Ԥ����ֵ

	//����������ֵ
	set._dAOutput1=MBCHAR2INT((unsigned char)recvParam[142]);//D/A���1
	set._dAOutput2=MBCHAR2INT((unsigned char)recvParam[143]);//D/A���2
	set._IO1=MBCHAR2INT((unsigned char)recvParam[144]);//I/O���1
	set._IO2=MBCHAR2INT((unsigned char)recvParam[145]);//I/O���2

	return true;
}

//========================������ȡ����ͻ������=============================
//recvParam��socketֱ�ӷ��ص��ַ���
//retParam�����صĽṹ��
bool parseReadFlashParam(char *recvParam,FlashData &set) {
	//��ʱ����
	int i=0;
	char temp[1024];
	char temp1[1024];
	int length[7];

	//��֤CRC
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

	//��֤���,�Խṹ�帳ֵ
	//��ö�ȡ���������ò����ַ���
//	char readParam[131];
//	for (int i=0;i<STRUCT2STRING;i++)
//	{
		//���յ����ַ�����ʽΪ��10�ֽ��豸���롢1�ֽ����������4�ֽڲ������ȡ�
		//131�ֽھ��������4�ֽ�crcУ��
//		readParam[i]=recvParam[i+15];
//	}
	
	//����������ֵ
	for(i=0;i<10;i++) {//��¼���
		temp1[i]=recvParam[i+15];
	}
	temp1[10]='\0';
	strcpy(set._jilubianhao,temp1);

	for(i=0;i<10;i++) {//���������
		temp1[i]=recvParam[i+25];
	}
	temp1[10]='\0';
	strcpy(set._chuanGanQiNumber,temp1);

	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+35];
	}
	temp[8]='\0';
	set._jingDu=MBCHAR2FLOAT(temp);//����
	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+43];
	}
	temp[8]='\0';
	set._weiDu=MBCHAR2FLOAT(temp);//γ��
	set._changDi=MBCHAR2INT((unsigned char)recvParam[51]);//��������
	set._guanCe=MBCHAR2INT((unsigned char)recvParam[52]);//�۲����
	set._anZhuangFangXiang=MBCHAR2INT((unsigned char)recvParam[53]);//��װ����

	//�����������
	for(i=0;i<6;i++) {
		temp1[i]=recvParam[54+i];
	}
	temp1[6]='\0';
	strcpy(set._xingZheng,temp1);

	//��ʼʱ��
	for(i=0;i<6;i++) {
		temp1[i]=(char)((MBCHAR2INT(recvParam[60+i*2]) << 4) + MBCHAR2INT(recvParam[60+i*2+1]));
	}
	temp1[6]='\0';
	strcpy(set._chushishike,temp1);

	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+72];
	}
	temp[8]='\0';
	set._caiYang=MBCHAR2FLOAT(temp);//��������
	
	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+80];
	}
	temp[8]='\0';
	set._pgazhi=MBCHAR2FLOAT(temp);//PGAֵ

	for(i=0;i<8;i++) {
		temp[i]=recvParam[i+88];
	}
	temp[8]='\0';
	set._siZhi=MBCHAR2FLOAT(temp);//SIֵ

	//��¼����
	for(i=0;i<7;i++) {
		length[i] = MBCHAR2INT((unsigned char)recvParam[i+96]);
	}
	set._jiluchangdu = length[6]*1000000+length[5]*100000+length[4]*10000+length[3]*1000+length[2]*100+length[1]*10+length[0];

	return true;
}

//����ɾ��Flash�д洢����ָ��
//���������
//recvParam��socketֱ�ӷ��ص��ַ���
bool ParseDelParam(char *recvParam) {
	int i;
	char result;
	char *temp1="";
	char str2[5];

	//��֤CRC
	for(i=0;i<(strlen(recvParam)-4);i++) {
		temp1[i]=recvParam[i];
	}
	
	sendStr(temp1,str2);

	strcat(temp1,str2);
	if(strcmp(recvParam,temp1)) {
		return false;
	}

	//��֤��ϣ����ؽ��
	result=recvParam[15];
	
	//0x01ʱ��ʾ���óɹ���0x00ʱ��ʾ����ʧ��
	//0x01��socket����ʱ��ת��Ϊ�ַ�'1'
	//0x00��socket����ʱ��ת��Ϊ�ַ�'0'
	if(!(result-'1')) {
		return true;
	} else {
		return false;
	}
}


//�������ò���ָ��ķ���ֵ
//���������
//recvParam��socketֱ�ӷ��ص��ַ���
bool ParseSetParam(char *recvParam) {
	char result;
	
	/*
	char *temp1="";
	char str2[5];
	int i;

	//��֤CRC
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
	//��֤��ϣ����ؽ��
	result=recvParam[15];

	//Э��涨��0x01ʱ��ʾ���óɹ���0x00ʱ��ʾ����ʧ��
	//0x01��socket����ʱ��ת��Ϊ�ַ�'1'
	//0x00��socket����ʱ��ת��Ϊ�ַ�'0'
	if(!(result-'1')) {
		//��resultΪ1ʱ
		return true;
	} else {
		return false;
	}
}


//��ȡ�洢��Flash�д�����������ָ��
//�ְ����ͣ����ĸ�ʽΪ��
//========================================================================================
//=֡ͷ���̶���0x00H������ַ���豸��ţ��������루�̶���'r'�������ݴ�С�����ݼ���õ�����=
//=�����ֽ����������Ҫ��������ݣ���CRCУ��λ��֡β���̶���0x7FH��                      =
//========================================================================================
//�������˽ӽ�������
//recvFlashData: ������������豸���յ����ַ������������ݼ����������ĸ�ʽ����
//ret: ���ز�����һ������������
bool parseRecvFlashData(char *recvFlashData, float *ret) {
	int i,temp;
	int recvDataLen=0;

	//�ж�֡ͷ�Ƿ�'00'
	if(recvFlashData[0]!='0' || recvFlashData[1]!='0') {
		return false;
	}
	//�жϹ������Ƿ�'r'
	if(recvFlashData[12]!='r') {
		return false;
	}

	//�������ݵĴ�С�Ƿ���ȷ
	for(i=13;i<21;i++) {
		temp=MBCHAR2INT((unsigned char)recvFlashData[i]);

		recvDataLen+=(temp<<((i-13)*4));
	}
	if((recvDataLen=+27) != strlen(recvFlashData)) {
		return false;
	}

	//CRCУ��
	

	//��ʼ����
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

//����Flash����
//deviceNum���豸���
//readData[]�����Ҷȼ��յ�������
//frameNum��֡��
//retString�����ɵ��ַ���
bool generateFlashData(char *deviceNum, short readData[480], short frameNum, char *retString) {
	char temp[1024];
	int i = 0;
	char str2[5];

	strcpy(temp, deviceNum);
	strcat(temp, "r01E0");

	for(i=0; i<480; i++) {
		//����ط��ĳ���д���ˣ��Ժ���ݲ�ͬ�Ĳ���Ƶ�ʿ��ܻ����޸�
		temp[15+i*2] = MBINT2CHAR((unsigned char)(readData[i]/16));
		temp[15+i*2+1] = MBINT2CHAR((unsigned char)(readData[i]%16));
	}
	
	temp[975] = '\0';
	
	//crcУ��
	sendStr(temp,str2);
	strcat(temp,str2);

	temp[979] = MBINT2CHAR((unsigned char)(frameNum/16));
	temp[980] = MBINT2CHAR((unsigned char)(frameNum%16));
	temp[981] = '\0';

	return true;
}

//����Flash����
//receiveStr���յ����ַ���
//readData[]��������������
//deviceID���豸���
//����ֵ��֡��
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