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
define("esri/dijit/PopupTemplate",["dojo/_base/declare","dojo/_base/lang","dojo/has","dojo/dom-construct","esri/kernel","esri/InfoTemplate","esri/PopupInfo","esri/dijit/PopupRenderer"],function(_1,_2,_3,_4,_5,_6,_7,_8){var PT=_1([_6,_7],{declaredClass:"esri.dijit.PopupTemplate","-chains-":{constructor:"manual"},chartTheme:null,constructor:function(_9,_a){_2.mixin(this,_a);this.initialize(_9,_a);},getTitle:function(_b){return this.info?this.getComponents(_b).title:"";},getContent:function(_c){return this.info?new _8({template:this,graphic:_c,chartTheme:this.chartTheme},_4.create("div")).domNode:"";}});if(_3("extend-esri")){_2.setObject("dijit.PopupTemplate",PT,_5);}return PT;});