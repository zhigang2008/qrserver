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
define("esri/tasks/SpatialRelationship",["dojo/_base/lang","dojo/has","esri/kernel"],function(_1,_2,_3){var _4={SPATIAL_REL_INTERSECTS:"esriSpatialRelIntersects",SPATIAL_REL_CONTAINS:"esriSpatialRelContains",SPATIAL_REL_CROSSES:"esriSpatialRelCrosses",SPATIAL_REL_ENVELOPEINTERSECTS:"esriSpatialRelEnvelopeIntersects",SPATIAL_REL_INDEXINTERSECTS:"esriSpatialRelIndexIntersects",SPATIAL_REL_OVERLAPS:"esriSpatialRelOverlaps",SPATIAL_REL_TOUCHES:"esriSpatialRelTouches",SPATIAL_REL_WITHIN:"esriSpatialRelWithin",SPATIAL_REL_RELATION:"esriSpatialRelRelation"};if(_2("extend-esri")){_1.setObject("tasks._SpatialRelationship",_4,_3);}return _4;});