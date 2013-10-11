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
define("esri/renderers/TemporalRenderer",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/renderers/Renderer"],function(_1,_2,_3,_4,_5){var _6=_1(_5,{declaredClass:"esri.renderer.TemporalRenderer",constructor:function(_7,_8,_9,_a){this.observationRenderer=_7;this.latestObservationRenderer=_8;this.trackRenderer=_9;this.observationAger=_a;},getSymbol:function(_b){var _c=_b.getLayer();var _d=_c._getKind(_b);var _e=(_d===0)?this.observationRenderer:(this.latestObservationRenderer||this.observationRenderer);var _f=(_e&&_e.getSymbol(_b));var _10=this.observationAger;if(_c.timeInfo&&_c._map.timeExtent&&(_e===this.observationRenderer)&&_10&&_f){_f=_10.getAgedSymbol(_f,_b);}return _f;}});if(_3("extend-esri")){_2.setObject("renderer.TemporalRenderer",_6,_4);}return _6;});