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
define("esri/dijit/editing/tools/Cut",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/_base/connect","dojo/has","dojo/DeferredList","esri/graphicsUtils","esri/graphic","esri/tasks/query","esri/layers/FeatureLayer","esri/toolbars/draw","esri/dijit/editing/Cut","esri/dijit/editing/tools/ToggleToolBase","esri/kernel"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9,_a,_b,_c,_d,_e){var _f=_1([_d],{declaredClass:"esri.dijit.editing.tools.Cut",id:"btnFeatureCut",_enabledIcon:"toolbarIcon cutIcon",_disabledIcon:"toolbarIcon cutIcon",_drawType:_b.POLYLINE,_enabled:true,_label:"NLS_cutLbl",_cutConnects:[],activate:function(){this._cutConnects.push(_4.connect(this._toolbar,"onDrawEnd",this,"_onDrawEnd"));this.inherited(arguments);},deactivate:function(){this.inherited(arguments);_3.forEach(this._cutConnects,_4.disconnect);this._cutConnects=[];this._edits=[];},_onDrawEnd:function(_10){var _11=this._settings.layers;var _12=this._cutLayers=_3.filter(_11,function(_13){return ((_13.geometryType==="esriGeometryPolygon")||(_13.geometryType==="esriGeometryPolyline")&&_13.visible&&_13._isMapAtVisibleScale());});this._cutConnects=this._cutConnects.concat(_3.map(_12,_2.hitch(this,function(_14){return _4.connect(_14,"onEditsComplete",_2.hitch(this,function(_15,_16,_17){if(this._settings.undoRedoManager){var _18=this._settings.undoRedoManager;_3.forEach(this._edits,_2.hitch(this,function(_19){_18.add(new _c({featureLayer:_19.layer,addedGraphics:_19.adds,preUpdatedGraphics:_19.preUpdates,postUpdatedGraphics:_19.updates}));}),this);}this.onFinished();}));})));var _1a=new _9();_1a.geometry=_10;_3.forEach(_12,function(_1b,idx){this._settings.editor._selectionHelper.selectFeatures([_1b],_1a,_a.SELECTION_NEW,_2.hitch(this,"_cutFeatures",_1b,_1a));},this);},_cutFeatures:function(_1c,_1d,_1e){if(!_1e||!_1e.length){return;}this._edits=[];var _1f=[];_1f.push(this._settings.geometryService.cut(_7.getGeometries(_1e),_1d.geometry,_2.hitch(this,"_cutHandler",_1c,_1e)));var _20=new _6(_1f).addCallback(_2.hitch(this,function(){this.onApplyEdits(this._edits);}));},_cutHandler:function(_21,_22,_23){var _24=[];var _25=[];var _26=_3.map(_22,function(_27){return new _8(_2.clone(_27.toJson()));});var _28;var _29;_3.forEach(_23.cutIndexes,function(_2a,i){if(_28!=_2a){_28=_2a;_25.push(_22[_2a].setGeometry(_23.geometries[i]));}else{_29=new _8(_23.geometries[i],null,_2.mixin({},_22[_2a].attributes),null);_29.attributes[_22[0].getLayer().objectIdField]=null;_24.push(_29);}},this);this._edits.push({layer:_21,adds:_24,updates:_25,preUpdates:_26});}});if(_5("extend-esri")){_2.setObject("dijit.editing.tools.Cut",_f,_e);}return _f;});