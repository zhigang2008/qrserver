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
define("esri/dijit/analysis/AnalysisToggleButton",["require","dojo/_base/declare","dojo/_base/lang","dojo/_base/connect","dojo/has","dojo/dom-class","dijit/_Widget","esri/kernel"],function(_1,_2,_3,_4,_5,_6,_7,_8){var _9=_2([_7],{groupName:"defaultGroup",declaredClass:"esri.dijit.analysis.AnalysisToggleButton",postMixInProperties:function(){this.inherited(arguments);this.unselectChannel="/ButtonGroupCtr/"+this.groupName;_4.subscribe(this.unselectChannel,this,"doUnselect");},postCreate:function(){this.inherited(arguments);_6.add(this.domNode,"esriGroupButton");},doUnselect:function(_a){if(_a!==this&&this.get("checked")){this.set("checked",false);}},_getCheckedAttr:function(){return this.checked;},_setCheckedAttr:function(_b){this.inherited(arguments);this.checked=_b;if(_b){_4.publish(this.unselectChannel,[this]);}_6.toggle(this.domNode,"esriGroupselected",_b);}});if(_5("extend-esri")){_3.setObject("dijit.analysis.AnalysisToggleButton",_9,_8);}return _9;});