package v1alpha1

import (
	prv1 "github.com/krateoplatformops/provider-runtime/apis/common/v1"
)

// GetCondition of this TeamRepo.
func (mg *TeamRepo) GetCondition(ct prv1.ConditionType) prv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this TeamRepo.
func (mg *TeamRepo) GetDeletionPolicy() prv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// SetConditions of this TeamRepo.
func (mg *TeamRepo) SetConditions(c ...prv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this TeamRepo.
func (mg *TeamRepo) SetDeletionPolicy(r prv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}
