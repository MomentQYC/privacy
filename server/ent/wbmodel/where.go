// Code generated by entc, DO NOT EDIT.

package wbmodel

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/kallydev/privacy/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// WBNumber applies equality check predicate on the "uid" field. It's identical to WBNumberEQ.
func WbNumber(v int64) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWbNumber), v))
	})
}

// PhoneNumber applies equality check predicate on the "phone_number" field. It's identical to PhoneNumberEQ.
func PhoneNumber(v int64) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhoneNumber), v))
	})
}

// WbNumberEQ applies the EQ predicate on the "uid" field.
func WbNumberEQ(v int64) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWbNumber), v))
	})
}

// WbNumberNEQ applies the NEQ predicate on the "uid" field.
func WbNumberNEQ(v int64) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWbNumber), v))
	})
}

// WbNumberIn applies the In predicate on the "uid" field.
func WbNumberIn(vs ...int64) predicate.WBModel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WBModel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWbNumber), v...))
	})
}

// WBNumberNotIn applies the NotIn predicate on the "uid" field.
func WbNumberNotIn(vs ...int64) predicate.WBModel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WBModel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWbNumber), v...))
	})
}

// WbNumberGT applies the GT predicate on the "uid" field.
func WbNumberGT(v int64) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWbNumber), v))
	})
}

// WbNumberGTE applies the GTE predicate on the "uid" field.
func WbNumberGTE(v int64) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWbNumber), v))
	})
}

// WbNumberLT applies the LT predicate on the "uid" field.
func WbNumberLT(v int64) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWbNumber), v))
	})
}

// WbNumberLTE applies the LTE predicate on the "uid" field.
func WbNumberLTE(v int64) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWbNumber), v))
	})
}

// PhoneNumberEQ applies the EQ predicate on the "phone_number" field.
func PhoneNumberEQ(v int64) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberNEQ applies the NEQ predicate on the "phone_number" field.
func PhoneNumberNEQ(v int64) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberIn applies the In predicate on the "phone_number" field.
func PhoneNumberIn(vs ...int64) predicate.WBModel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WBModel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPhoneNumber), v...))
	})
}

// PhoneNumberNotIn applies the NotIn predicate on the "phone_number" field.
func PhoneNumberNotIn(vs ...int64) predicate.WBModel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WBModel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPhoneNumber), v...))
	})
}

// PhoneNumberGT applies the GT predicate on the "phone_number" field.
func PhoneNumberGT(v int64) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberGTE applies the GTE predicate on the "phone_number" field.
func PhoneNumberGTE(v int64) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberLT applies the LT predicate on the "phone_number" field.
func PhoneNumberLT(v int64) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberLTE applies the LTE predicate on the "phone_number" field.
func PhoneNumberLTE(v int64) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPhoneNumber), v))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.WBModel) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.WBModel) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WBModel) predicate.WBModel {
	return predicate.WBModel(func(s *sql.Selector) {
		p(s.Not())
	})
}
