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
require({cache:{"url:esri/dijit/analysis/templates/AnalysisToolItem.html":"<div class='toolContainer' data-dojo-attach-point=\"_toolCtr\" style=\"cursor:pointer;cursor:hand;\" data-dojo-attach-event=\"onclick:_handleToolIconClick\">\r\n  <div data-dojo-attach-point='_toolIcon'></div>\r\n  <div class='esriLeadingMargin5' style='margin-top:-42px;'>\r\n    <a  href=\"#\" class='esriFloatTrailing helpIcon' esriHelpTopic=\"toolDescription\" data-dojo-attach-point=\"_helpIconNode\"></a>\r\n  \t<label data-dojo-attach-point='_toolNameLabel' style=\"cursor:pointer;cursor:hand;\"></label>\r\n  </div>\r\n  <div class='esriLeadingMargin2' data-dojo-attach-point=\"optionsDiv\" style=\"margin-top:0.5em;font-size:0.85em;\"><label class=\"esriLeadingMargin5 comingSoonIcon\">${i18n.comingSoonLabel}</label></div>\t\r\n</div>\r\n"}});define("esri/dijit/analysis/AnalysisToolItem",["require","dojo/_base/declare","dojo/_base/lang","dojo/_base/connect","dojo/_base/event","dojo/has","dojo/dom-class","dojo/dom-attr","dojo/dom-style","dijit/_WidgetBase","dijit/_TemplatedMixin","dijit/_OnDijitClickMixin","dijit/_FocusMixin","esri/kernel","dojo/text!esri/dijit/analysis/templates/AnalysisToolItem.html"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9,_a,_b,_c,_d,_e,_f){var _10=_2([_a,_b,_c,_d],{declaredClass:"esri.dijit.analysis.AnalysisToolItem",templateString:_f,basePath:_1.toUrl("esri/dijit/analysis"),widgetsInTemplate:true,i18n:null,_helpIconNode:null,_toolIcon:null,_toolIconClass:null,_toolNameLabel:null,toolName:null,helpTopic:null,helpFileName:"Analysis",constructor:function(_11,_12){this.inherited(arguments);if(_11.toolIcon){this._toolIconClass=_11.toolIcon;}if(_11.name){this.toolName=_11.name;this.helpTopic=_11.helpTopic;}},postCreate:function(){this.inherited(arguments);this._toolNameLabel.innerHTML=this.toolName;_7.add(this._toolIcon,this._toolIconClass);_8.set(this._helpIconNode,"esriHelpTopic",this.helpTopic);this.set("showComingSoonLabel",true);},postMixInProperties:function(){this.inherited(arguments);this.i18n={};_3.mixin(this.i18n,dojo.i18n.getLocalization("esri","jsapi").common);_3.mixin(this.i18n,dojo.i18n.getLocalization("esri","jsapi").analysisTools);},_handleToolNameClick:function(){this.onToolSelect(this);},_handleToolIconClick:function(e){_5.stop(e);this.onToolSelect(this);},_setShowComingSoonLabelAttr:function(_13){_9.set(this.optionsDiv,"display",(_13===true)?"block":"none");_7.toggle(this._toolCtr,"esriToolContainerDisabled",_13);_7.toggle(this._toolNameLabel,"esriTransparentNode",_13);_7.toggle(this._toolIcon,"esriTransparentNode",_13);_9.set(this._toolNameLabel,"cursor",(_13===true)?"default":"pointer");_9.set(this._toolCtr,"cursor",(_13===true)?"default":"pointer");},onToolSelect:function(_14){}});if(_6("extend-esri")){_3.setObject("dijit.analysis.AnalysisToolItem",_10,_e);}return _10;});