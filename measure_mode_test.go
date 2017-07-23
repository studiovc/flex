package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type _MeasureConstraint struct {
	width      float32
	widthMode  YGMeasureMode
	height     float32
	heightMode YGMeasureMode
}

type _MeasureConstraintList struct {
	length      int
	constraints []_MeasureConstraint
}

func _measure2(node *YGNode,
	width float32,
	widthMode YGMeasureMode,
	height float32,
	heightMode YGMeasureMode) YGSize {
	constraintList := YGNodeGetContext(node).(*_MeasureConstraintList)
	constraints := constraintList.constraints
	currentIndex := constraintList.length
	(&constraints[currentIndex]).width = width
	(&constraints[currentIndex]).widthMode = widthMode
	(&constraints[currentIndex]).height = height
	(&constraints[currentIndex]).heightMode = heightMode
	constraintList.length = currentIndex + 1

	if widthMode == YGMeasureModeUndefined {
		width = 10
	}

	if heightMode == YGMeasureModeUndefined {
		height = 10
	} else {
		height = width // TODO:: is it a bug in tests ?
	}
	return YGSize{
		width:  width,
		height: height,
	}
}

func TestExactly_measure_stretched_child_column(t *testing.T) {
	constraintList := _MeasureConstraintList{
		length:      0,
		constraints: make([]_MeasureConstraint, 10, 10),
	}

	root := YGNodeNew()
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNew()
	YGNodeSetContext(rootChild0, &constraintList)
	YGNodeSetMeasureFunc(rootChild0, _measure2)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assert.Equal(t, 1, constraintList.length)

	assertFloatEqual(t, 100, constraintList.constraints[0].width)
	assert.Equal(t, YGMeasureModeExactly, constraintList.constraints[0].widthMode)

}

func TestExactly_measure_stretched_child_row(t *testing.T) {
	constraintList := _MeasureConstraintList{
		length:      0,
		constraints: make([]_MeasureConstraint, 10, 10),
	}

	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNew()
	YGNodeSetContext(rootChild0, &constraintList)
	YGNodeSetMeasureFunc(rootChild0, _measure2)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assert.Equal(t, 1, constraintList.length)

	assertFloatEqual(t, 100, constraintList.constraints[0].height)
	assert.Equal(t, YGMeasureModeExactly, constraintList.constraints[0].heightMode)

}

func TestAt_most_main_axis_column(t *testing.T) {
	constraintList := _MeasureConstraintList{
		length:      0,
		constraints: make([]_MeasureConstraint, 10, 10),
	}

	root := YGNodeNew()
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNew()
	YGNodeSetContext(rootChild0, &constraintList)
	YGNodeSetMeasureFunc(rootChild0, _measure2)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assert.Equal(t, 1, constraintList.length)

	assertFloatEqual(t, 100, constraintList.constraints[0].height)
	assert.Equal(t, YGMeasureModeAtMost, constraintList.constraints[0].heightMode)

}

func TestAt_most_cross_axis_column(t *testing.T) {
	constraintList := _MeasureConstraintList{
		length:      0,
		constraints: make([]_MeasureConstraint, 10, 10),
	}

	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNew()
	YGNodeSetContext(rootChild0, &constraintList)
	YGNodeSetMeasureFunc(rootChild0, _measure2)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assert.Equal(t, 1, constraintList.length)

	assertFloatEqual(t, 100, constraintList.constraints[0].width)
	assert.Equal(t, YGMeasureModeAtMost, constraintList.constraints[0].widthMode)

}

func TestAt_most_main_axis_row(t *testing.T) {
	constraintList := _MeasureConstraintList{
		length:      0,
		constraints: make([]_MeasureConstraint, 10, 10),
	}

	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNew()
	YGNodeSetContext(rootChild0, &constraintList)
	YGNodeSetMeasureFunc(rootChild0, _measure2)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assert.Equal(t, 1, constraintList.length)

	assertFloatEqual(t, 100, constraintList.constraints[0].width)
	assert.Equal(t, YGMeasureModeAtMost, constraintList.constraints[0].widthMode)

}

func TestAt_most_cross_axis_row(t *testing.T) {
	constraintList := _MeasureConstraintList{
		length:      0,
		constraints: make([]_MeasureConstraint, 10, 10),
	}

	root := YGNodeNew()
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetWidth(root, 100)
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNew()
	YGNodeSetContext(rootChild0, &constraintList)
	YGNodeSetMeasureFunc(rootChild0, _measure2)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assert.Equal(t, 1, constraintList.length)

	assertFloatEqual(t, 100, constraintList.constraints[0].height)
	assert.Equal(t, YGMeasureModeAtMost, constraintList.constraints[0].heightMode)

}

func TestFlex_child(t *testing.T) {
	constraintList := _MeasureConstraintList{
		length:      0,
		constraints: make([]_MeasureConstraint, 10, 10),
	}

	root := YGNodeNew()
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNew()
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeSetContext(rootChild0, &constraintList)
	YGNodeSetMeasureFunc(rootChild0, _measure2)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assert.Equal(t, 2, constraintList.length)

	assertFloatEqual(t, 100, constraintList.constraints[0].height)
	assert.Equal(t, YGMeasureModeAtMost, constraintList.constraints[0].heightMode)

	assertFloatEqual(t, 100, constraintList.constraints[1].height)
	assert.Equal(t, YGMeasureModeExactly, constraintList.constraints[1].heightMode)

}

func TestFlex_child_with_flex_basis(t *testing.T) {
	constraintList := _MeasureConstraintList{
		length:      0,
		constraints: make([]_MeasureConstraint, 10, 10),
	}

	root := YGNodeNew()
	YGNodeStyleSetHeight(root, 100)

	rootChild0 := YGNodeNew()
	YGNodeStyleSetFlexGrow(rootChild0, 1)
	YGNodeStyleSetFlexBasis(rootChild0, 0)
	YGNodeSetContext(rootChild0, &constraintList)
	YGNodeSetMeasureFunc(rootChild0, _measure2)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assert.Equal(t, 1, constraintList.length)

	assertFloatEqual(t, 100, constraintList.constraints[0].height)
	assert.Equal(t, YGMeasureModeExactly, constraintList.constraints[0].heightMode)

}

func TestOverflow_scroll_column(t *testing.T) {
	constraintList := _MeasureConstraintList{
		length:      0,
		constraints: make([]_MeasureConstraint, 10, 10),
	}

	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetOverflow(root, YGOverflowScroll)
	YGNodeStyleSetHeight(root, 100)
	YGNodeStyleSetWidth(root, 100)

	rootChild0 := YGNodeNew()
	YGNodeSetContext(rootChild0, &constraintList)
	YGNodeSetMeasureFunc(rootChild0, _measure2)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assert.Equal(t, 1, constraintList.length)

	assertFloatEqual(t, 100, constraintList.constraints[0].width)
	assert.Equal(t, YGMeasureModeAtMost, constraintList.constraints[0].widthMode)

	assert.True(t, YGFloatIsUndefined(constraintList.constraints[0].height))
	assert.Equal(t, YGMeasureModeUndefined, constraintList.constraints[0].heightMode)

}

func TestOverflow_scroll_row(t *testing.T) {
	constraintList := _MeasureConstraintList{
		length:      0,
		constraints: make([]_MeasureConstraint, 10, 10),
	}

	root := YGNodeNew()
	YGNodeStyleSetAlignItems(root, AlignFlexStart)
	YGNodeStyleSetFlexDirection(root, YGFlexDirectionRow)
	YGNodeStyleSetOverflow(root, YGOverflowScroll)
	YGNodeStyleSetHeight(root, 100)
	YGNodeStyleSetWidth(root, 100)

	rootChild0 := YGNodeNew()
	YGNodeSetContext(rootChild0, &constraintList)
	YGNodeSetMeasureFunc(rootChild0, _measure2)
	YGNodeInsertChild(root, rootChild0, 0)

	YGNodeCalculateLayout(root, YGUndefined, YGUndefined, DirectionLTR)

	assert.Equal(t, 1, constraintList.length)

	assert.True(t, YGFloatIsUndefined(constraintList.constraints[0].width))
	assert.Equal(t, YGMeasureModeUndefined, constraintList.constraints[0].widthMode)

	assertFloatEqual(t, 100, constraintList.constraints[0].height)
	assert.Equal(t, YGMeasureModeAtMost, constraintList.constraints[0].heightMode)

}