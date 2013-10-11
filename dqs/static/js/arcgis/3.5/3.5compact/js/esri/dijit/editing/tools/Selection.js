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
define("esri/dijit/editing/tools/Selection",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/_base/connect","dojo/_base/Color","dojo/has","esri/symbols/SimpleMarkerSymbol","esri/symbols/SimpleLineSymbol","esri/symbols/SimpleFillSymbol","esri/dijit/editing/tools/Edit","esri/dijit/editing/tools/SelectionTools","esri/dijit/editing/tools/DropDownToolBase","esri/kernel"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9,_a,_b,_c,_d){var _e=_1([_c],{declaredClass:"esri.dijit.editing.tools.Selection",_enabled:true,activate:function(){this.inherited(arguments);this._sConnect=_4.connect(this._toolbar,"onDrawEnd",this,"_onDrawEnd");},deactivate:function(){this.inherited(arguments);_4.disconnect(this._sConnect);delete this._sConnect;},_initializeTool:function(){this._createSymbols();this._initializeLayers();this._toolTypes=["select","selectadd","selectremove"];},_onDrawEnd:function(_f){this.inherited(arguments);this._settings.editor._hideAttributeInspector();var _10=this._settings.layers;this._selectMethod=this._activeTool._selectMethod;this._settings.editor._selectionHelper.selectFeaturesByGeometry(_10,_f,this._selectMethod,_2.hitch(this,"onFinished"));},_createSymbols:function(){this._pointSelectionSymbol=new _7(_7.STYLE_CIRCLE,10,new _8(_8.STYLE_SOLID,new _5([0,0,0]),1),new _5([255,0,0,0.5]));this._polylineSelectionSymbol=new _8(_8.STYLE_SOLID,new _5([0,200,255]),2);this._polygonSelectionSymbol=new _9(_9.STYLE_SOLID,new _8(_8.STYLE_SOLID,new _5([0,0,0]),1),new _5([0,200,255,0.5]));},_initializeLayers:function(){var _11=this._settings.layers;_3.forEach(_11,this._setSelectionSymbol,this);},_setSelectionSymbol:function(_12){var _13=null;switch(_12.geometryType){case "esriGeometryPoint":_13=this._pointSelectionSymbol;break;case "esriGeometryPolyline":_13=this._polylineSelectionSymbol;break;case "esriGeometryPolygon":_13=this._polygonSelectionSymbol;break;}_12.setSelectionSymbol(_12._selectionSymbol||_13);},_createTools:function(){_3.forEach(this._toolTypes,this._createTool,this);this.inherited(arguments);},_createTool:function(_14){var _15=_2.mixin(_b[_14],{settings:this._settings,onClick:_2.hitch(this,"onItemClicked",_b[_14].id)});this._tools[_14.toUpperCase()]=new _a(_15);}});if(_6("extend-esri")){_2.setObject("dijit.editing.tools.Selection",_e,_d);}return _e;});