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
define("esri/dijit/editing/tools/Editing",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/has","esri/layers/FeatureTemplate","esri/dijit/editing/tools/Edit","esri/dijit/editing/tools/EditingTools","esri/dijit/editing/tools/DropDownToolBase","esri/kernel","esri/lang"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9,_a){var _b=_1([_8],{declaredClass:"esri.dijit.editing.tools.Editing",_enabled:false,deactivate:function(){if(!this._enabled){return;}this._enabled=false;this.inherited(arguments);this._settings.templatePicker.clearSelection();},onItemClicked:function(_c,_d){this.inherited(arguments);if(this._activeTool===this._tools.AUTOCOMPLETE){this._settings.editor._drawingTool=_5.TOOL_AUTO_COMPLETE_POLYGON;}},_activateTool:function(_e,_f){var i;this.enable(_f);for(i in this._tools){if(this._tools.hasOwnProperty(i)){this.dropDown.removeChild(this._tools[i]);if(this._tools[i]._enabled===true){this.dropDown.addChild(this._tools[i]);}}}if(this._geometryType!==_f||this._activeTool._enabled===false){this._activeTool=this._tools[_e.toUpperCase()];}this._geometryType=_f;this._activeTool.activate();this._activeTool.setChecked(true);this._updateUI();},_initializeTool:function(_10){this.inherited(arguments);this._initializeTools();},_initializeTools:function(){var _11=this._settings.layers;var _12=this._settings.editor;var _13=false,_14=false,_15=false;var _16=this._toolTypes=[];var _17;_3.forEach(_11,function(_18){_17=_18.geometryType;_13=_13||_17==="esriGeometryPoint";_14=_14||_17==="esriGeometryPolyline";_15=_15||_17==="esriGeometryPolygon";_16=_16.concat(_3.map(this._getTemplatesFromLayer(_18),_2.hitch(this,function(_19){return _12._toDrawTool(_19.drawingTool,_18);})));},this);var _1a=this._settings.createOptions;if(_13){this._toolTypes.push("point");}if(_14){this._toolTypes=this._toolTypes.concat(_1a.polylineDrawTools);}if(_15){this._toolTypes=this._toolTypes.concat(_1a.polygonDrawTools);}this._toolTypes=this._toUnique(this._toolTypes.concat(_16));},_toUnique:function(arr){var _1b={};return _3.filter(arr,function(val){return _1b[val]?false:(_1b[val]=true);});},_getTemplatesFromLayer:function(_1c){var _1d=_1c.templates||[];var _1e=_1c.types;_3.forEach(_1e,function(_1f){_1d=_1d.concat(_1f.templates);});return _3.filter(_1d,_a.isDefined);},_createTools:function(){_3.forEach(this._toolTypes,this._createTool,this);this.inherited(arguments);},_createTool:function(_20){var _21=_2.mixin(_7[_20],{settings:this._settings,onClick:_2.hitch(this,"onItemClicked",_7[_20].id)});this._tools[_20.toUpperCase()]=new _6(_21);}});if(_4("extend-esri")){_2.setObject("dijit.editing.tools.Editing",_b,_9);}return _b;});