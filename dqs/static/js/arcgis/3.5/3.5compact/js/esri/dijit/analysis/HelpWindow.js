/*
 COPYRIGHT 2009 ESRI

 TRADE SECRETS: ESRI PROPRIETARY AND CONFIDENTIAL
 Unpublished material - all rights reserved under the
 Copyright Laws of the United States and applicable international
 laws, treaties, and conventions.

 For additional information, contact:
 Environmental Systems Research Institute, Inc.
 Attn: Contracts and Legal Services Department
 380 New York Street
 Redlands, California, 92373
 USA

 email: contracts@esri.com
 */
//>>built
define("esri/dijit/analysis/HelpWindow",["require","dojo/_base/declare","dojo/_base/lang","dojo/_base/connect","dojo/_base/event","dojo/_base/kernel","dojo/has","dojo/i18n","dojo/dom-construct","dojo/dom-class","dojo/dom-attr","dojo/dom-style","dijit/_Widget","dijit/TooltipDialog","dijit/popup","esri/kernel"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9,_a,_b,_c,_d,_e,_f,_10){var _11=_2([_d],{declaredClass:"esri.dijit.analysis.HelpWindow",i18n:null,basePath:_1.toUrl("esri/dijit/analysis"),postMixInProperties:function(){this.inherited(arguments);this.i18n={};_3.mixin(this.i18n,_8.getLocalization("esri","jsapi").common);_3.mixin(this.i18n,_8.getLocalization("esri","jsapi").analysisHelp);},postCreate:function(){this.inherited(arguments);},_computeSize:function(_12){var _13={w:400,h:200};if(_12.indexOf("Category")!==-1){_13.w=400;_13.h=320;}else{if(_12.indexOf("Tool")!==-1){_13.w=400;_13.h=320;}else{if(_12.indexOf("toolDescription")!==-1){_13.w=400;_13.h=520;}}}return _13;},_setHelpTopicAttr:function(_14){if(this.tooltipHelpDlg){_f.close(this.tooltipHelpDlg);_9.destroy(this.tooltipHelpDlg);_9.destroy(this.tooltipHelpDlg.domNode);}var _15=_6.baseUrl.substring(0,_6.baseUrl.indexOf("/js/"));var _16=this._computeSize(_14);var _17="<div class='' style='position=relative'"+"<div class='sizer content'>"+"<div class='contentPane'>"+"<div class='esriFloatTrailing' style='padding:0;'>"+"<a href='#' onclick='esri.dijit._helpDialog.close()' 'title='"+this.i18n.close+"'>"+"<img src='images/close.gif' border='0'/></a>"+"</div>"+"<iframe frameborder='0'  id='"+_14+"' src='"+_15+"/js/esri/dijit/analysis/help/"+this.helpFileName+".html#"+_14+"' width='"+_16.w+"' height='"+_16.h+"' marginheight='0' marginwidth='0'></iframe>"+"</div>"+"</div>"+"<div class='sizer'>"+"<div class='actionsPane'>"+"<div class='actionList hidden'>"+"<a class='action zoomTo' href='"+_15+"/js/esri/dijit/analysis/help/"+this.helpFileName+".html' target='_help'>"+"Learn More"+"</a>"+"</div>"+"</div>"+"</div>"+"</div>"+"</div>";this.tooltipHelpDlg=new _e({"preload":true,"content":_17,"class":"esriHelpPopup esriHelpPopupWrapper"});},show:function(_18,_19,_1a){this.helpFileName=_1a;this.set("helpTopic",_19);_f.open({popup:this.tooltipHelpDlg,x:_18.pageX+40,y:_18.screenY-_18.pageY+10,onCancel:function(){log(self.id+": cancel of child");},onExecute:function(){log(self.id+": execute of child");_f.close(this.tooltipHelpDlg);self.open=false;}});},close:function(_1b,_1c){_f.close(this.tooltipHelpDlg);}});if(_7("extend-esri")){_3.setObject("dijit.analysis.HelpWindow",_11,_10);}return _11;});