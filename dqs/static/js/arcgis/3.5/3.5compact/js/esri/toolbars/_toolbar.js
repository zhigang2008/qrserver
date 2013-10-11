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
define("esri/toolbars/_toolbar",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel"],function(_1,_2,_3,_4){var TB=_1(null,{declaredClass:"esri.toolbars._Toolbar",constructor:function(_5){this.map=_5;},_cursors:{"move":"pointer","move-v":"pointer","move-gv":"pointer","box0":"nw-resize","box1":"n-resize","box2":"ne-resize","box3":"e-resize","box4":"se-resize","box5":"s-resize","box6":"sw-resize","box7":"w-resize","box8":"pointer"},_deactivateMapTools:function(_6,_7,_8,_9){var _a=this.map;if(_6){this._mapNavState={isDoubleClickZoom:_a.isDoubleClickZoom,isClickRecenter:_a.isClickRecenter,isPan:_a.isPan,isRubberBandZoom:_a.isRubberBandZoom,isKeyboardNavigation:_a.isKeyboardNavigation,isScrollWheelZoom:_a.isScrollWheelZoom};_a.disableDoubleClickZoom();_a.disableClickRecenter();_a.disablePan();_a.disableRubberBandZoom();_a.disableKeyboardNavigation();}if(_7){_a.hideZoomSlider();}if(_8){_a.hidePanArrows();}if(_9){_a.graphics.disableMouseEvents();}},_activateMapTools:function(_b,_c,_d,_e){var _f=this.map,_10=this._mapNavState;if(_b&&_10){if(_10.isDoubleClickZoom){_f.enableDoubleClickZoom();}if(_10.isClickRecenter){_f.enableClickRecenter();}if(_10.isPan){_f.enablePan();}if(_10.isRubberBandZoom){_f.enableRubberBandZoom();}if(_10.isKeyboardNavigation){_f.enableKeyboardNavigation();}if(_10.isScrollWheelZoom){_f.enableScrollWheelZoom();}}if(_c){_f.showZoomSlider();}if(_d){_f.showPanArrows();}if(_e){_f.graphics.enableMouseEvents();}}});if(_3("extend-esri")){_2.setObject("toolbars._Toolbar",TB,_4);}return TB;});