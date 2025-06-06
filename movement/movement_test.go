package movement_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/PaloAltoNetworks/pango/v2/movement"
)

var _ = fmt.Printf

type Mock struct {
	Name string
}

func (o Mock) EntryName() string {
	return o.Name
}

func asMovable(mocks []string) []movement.Movable {
	var movables []movement.Movable

	for _, elt := range mocks {
		movables = append(movables, Mock{elt})
	}

	return movables
}

var _ = Describe("MoveGroup()", func() {
	Context("With PositionFirst used as position", func() {
		Context("when existing positions matches expected", func() {
			It("should generate no movements", func() {
				// '(A B C) -> '(A B C)
				expected := asMovable([]string{"A", "B", "C"})
				moves, err := movement.MoveGroup(movement.PositionFirst{}, expected, expected)
				Expect(err).ToNot(HaveOccurred())
				Expect(moves).To(HaveLen(0))
			})
		})
		Context("when it has to move two elements", func() {
			It("should generate three move actions", func() {
				// '(D E A B C) -> '(A B C D E)
				entries := asMovable([]string{"A", "B", "C"})
				existing := asMovable([]string{"D", "E", "A", "B", "C"})

				moves, err := movement.MoveGroup(movement.PositionFirst{}, entries, existing)
				Expect(err).ToNot(HaveOccurred())
				Expect(moves).To(HaveLen(3))

				Expect(moves[0].Movable.EntryName()).To(Equal("A"))
				Expect(moves[0].Where).To(Equal(movement.ActionWhereFirst))
				Expect(moves[0].Destination).To(BeNil())

				Expect(moves[1].Movable.EntryName()).To(Equal("B"))
				Expect(moves[1].Where).To(Equal(movement.ActionWhereAfter))
				Expect(moves[1].Destination.EntryName()).To(Equal("A"))

				Expect(moves[2].Movable.EntryName()).To(Equal("C"))
				Expect(moves[2].Where).To(Equal(movement.ActionWhereAfter))
				Expect(moves[2].Destination.EntryName()).To(Equal("B"))
			})
		})
		Context("when expected order is reversed", func() {
			It("should generate required move actions to converge lists", func() {
				// '(A B C D E) -> '(E D C B A)
				entries := asMovable([]string{"E", "D", "C", "B", "A"})
				existing := asMovable([]string{"A", "B", "C", "D", "E"})
				moves, err := movement.MoveGroup(movement.PositionFirst{}, entries, existing)
				Expect(err).ToNot(HaveOccurred())

				// '((E 'top nil)(B 'after E)(C 'after B)(D 'after C))
				// 'A element stays in place
				Expect(moves).To(HaveLen(4))
			})
		})
	})
	Context("With PositionLast used as position", func() {
		Context("with non-consecutive entries", func() {
			It("should generate two move actions", func() {
				// '(A E B C D) -> '(A B D E C)
				entries := asMovable([]string{"E", "C"})
				existing := asMovable([]string{"A", "E", "B", "C", "D"})

				moves, err := movement.MoveGroup(movement.PositionLast{}, entries, existing)
				Expect(err).ToNot(HaveOccurred())
				Expect(moves).To(HaveLen(2))

				Expect(moves[0].Movable.EntryName()).To(Equal("E"))
				Expect(moves[0].Where).To(Equal(movement.ActionWhereLast))
				Expect(moves[0].Destination).To(BeNil())

				Expect(moves[1].Movable.EntryName()).To(Equal("C"))
				Expect(moves[1].Where).To(Equal(movement.ActionWhereAfter))
				Expect(moves[1].Destination.EntryName()).To(Equal("E"))
			})
		})
	})
	Context("With PositionLast used as position", func() {
		Context("when it needs to move one element", func() {
			It("should generate a single move action", func() {
				// '(A E B C D) -> '(A B C D E)
				entries := asMovable([]string{"E"})
				existing := asMovable([]string{"A", "E", "B", "C", "D"})

				moves, err := movement.MoveGroup(movement.PositionLast{}, entries, existing)
				Expect(err).ToNot(HaveOccurred())
				Expect(moves).To(HaveLen(1))

				Expect(moves[0].Movable.EntryName()).To(Equal("E"))
				Expect(moves[0].Where).To(Equal(movement.ActionWhereLast))
				Expect(moves[0].Destination).To(BeNil())
			})
		})
	})

	Context("With PositionAfter used as position", func() {
		existing := asMovable([]string{"A", "B", "C", "D", "E"})
		Context("when direct position relative to the pivot is not required", func() {
			It("should not generate any move actions", func() {
				// '(A B C D E) -> '(A B C D E)
				entries := asMovable([]string{"D", "E"})
				moves, err := movement.MoveGroup(
					movement.PositionAfter{Directly: false, Pivot: "B"},
					entries, existing,
				)

				Expect(err).ToNot(HaveOccurred())
				Expect(moves).To(HaveLen(0))
			})
			Context("and moved entries are out of order", func() {
				It("should generate a single command to move B before D", func() {
					// '(A B C D E) -> '(A B C E D)
					entries := asMovable([]string{"E", "D"})
					moves, err := movement.MoveGroup(
						movement.PositionAfter{Directly: false, Pivot: "B"},
						entries, existing,
					)

					Expect(err).ToNot(HaveOccurred())
					Expect(moves).To(HaveLen(1))

					Expect(moves[0].Movable.EntryName()).To(Equal("E"))
					Expect(moves[0].Where).To(Equal(movement.ActionWhereAfter))
					Expect(moves[0].Destination.EntryName()).To(Equal("C"))
				})
			})
		})
		Context("when direct position relative to the pivot is required", func() {
			It("should generate required move actions", func() {
				// '(A B C D E) -> '(C D A B E)
				entries := asMovable([]string{"A", "B"})
				moves, err := movement.MoveGroup(
					movement.PositionAfter{Directly: true, Pivot: "D"},
					entries, existing,
				)

				Expect(err).ToNot(HaveOccurred())
				Expect(moves).To(HaveLen(2))

				Expect(moves[0].Movable.EntryName()).To(Equal("A"))
				Expect(moves[0].Where).To(Equal(movement.ActionWhereAfter))
				Expect(moves[0].Destination.EntryName()).To(Equal("D"))

				Expect(moves[1].Movable.EntryName()).To(Equal("B"))
				Expect(moves[1].Where).To(Equal(movement.ActionWhereAfter))
				Expect(moves[1].Destination.EntryName()).To(Equal("A"))
			})
			It("should not create a list with null entries", func() {
				// '(A B C D E) -> '(A C D B E)
				entries := asMovable([]string{"C", "D"})
				moves, err := movement.MoveGroup(
					movement.PositionAfter{Directly: true, Pivot: "A"},
					entries, existing,
				)

				Expect(err).ToNot(HaveOccurred())
				Expect(moves).To(HaveLen(2))

				Expect(moves[0].Movable.EntryName()).To(Equal("C"))
				Expect(moves[0].Where).To(Equal(movement.ActionWhereAfter))
				Expect(moves[0].Destination.EntryName()).To(Equal("A"))

				Expect(moves[1].Movable.EntryName()).To(Equal("D"))
				Expect(moves[1].Where).To(Equal(movement.ActionWhereAfter))
				Expect(moves[1].Destination.EntryName()).To(Equal("C"))
			})
		})
		Context("when direct position relative to the pivot is required, and order changes", func() {
			It("should generate required move actions when after is used", func() {
				// '(A B C D E) -> '(C D B A E)
				entries := asMovable([]string{"B", "A"})
				moves, err := movement.MoveGroup(
					movement.PositionAfter{Directly: true, Pivot: "D"},
					entries, existing,
				)

				Expect(err).ToNot(HaveOccurred())
				Expect(moves).To(HaveLen(2))

				Expect(moves[0].Movable.EntryName()).To(Equal("B"))
				Expect(moves[0].Where).To(Equal(movement.ActionWhereAfter))
				Expect(moves[0].Destination.EntryName()).To(Equal("D"))

				Expect(moves[1].Movable.EntryName()).To(Equal("A"))
				Expect(moves[1].Where).To(Equal(movement.ActionWhereAfter))
				Expect(moves[1].Destination.EntryName()).To(Equal("B"))
			})
		})
		Context("when direct position relative to the pivot is required, and order stays the same", func() {
			It("should generate required move actions when after is used", func() {
				// '(A B C D E) -> '(C D A B E)
				entries := asMovable([]string{"A", "B"})
				moves, err := movement.MoveGroup(
					movement.PositionAfter{Directly: true, Pivot: "D"},
					entries, existing,
				)

				Expect(err).ToNot(HaveOccurred())
				Expect(moves).To(HaveLen(2))

				Expect(moves[0].Movable.EntryName()).To(Equal("A"))
				Expect(moves[0].Where).To(Equal(movement.ActionWhereAfter))
				Expect(moves[0].Destination.EntryName()).To(Equal("D"))

				Expect(moves[1].Movable.EntryName()).To(Equal("B"))
				Expect(moves[1].Where).To(Equal(movement.ActionWhereAfter))
				Expect(moves[1].Destination.EntryName()).To(Equal("A"))
			})
		})
		Context("when direct position relative to the pivot is required, and order changes", func() {
			It("should generate required move actions when before is used", func() {
				// '(A B C D E) -> '(C D B A E)
				entries := asMovable([]string{"B", "A"})
				moves, err := movement.MoveGroup(
					movement.PositionBefore{Directly: true, Pivot: "E"},
					entries, existing,
				)

				Expect(err).ToNot(HaveOccurred())
				Expect(moves).To(HaveLen(2))

				Expect(moves[0].Movable.EntryName()).To(Equal("B"))
				Expect(moves[0].Where).To(Equal(movement.ActionWhereBefore))
				Expect(moves[0].Destination.EntryName()).To(Equal("E"))

				Expect(moves[1].Movable.EntryName()).To(Equal("A"))
				Expect(moves[1].Where).To(Equal(movement.ActionWhereAfter))
				Expect(moves[1].Destination.EntryName()).To(Equal("B"))
			})
		})
		Context("when direct position relative to the pivot is required, and order stays the same", func() {
			It("should generate required move actions when before is used", func() {
				// '(A B C D E) -> '(C D A B E)
				entries := asMovable([]string{"A", "B"})
				moves, err := movement.MoveGroup(
					movement.PositionBefore{Directly: true, Pivot: "E"},
					entries, existing,
				)

				Expect(err).ToNot(HaveOccurred())
				Expect(moves).To(HaveLen(2))

				Expect(moves[0].Movable.EntryName()).To(Equal("A"))
				Expect(moves[0].Where).To(Equal(movement.ActionWhereBefore))
				Expect(moves[0].Destination.EntryName()).To(Equal("E"))

				Expect(moves[1].Movable.EntryName()).To(Equal("B"))
				Expect(moves[1].Where).To(Equal(movement.ActionWhereAfter))
				Expect(moves[1].Destination.EntryName()).To(Equal("A"))
			})
		})
	})

	// '(A E B C D) -> '(A B C D E) => '(E 'bottom nil) / '(E 'after D)

	// PositionSomewhereBefore PositionDirectlyBefore
	// '(C B 'before E, directly)
	// '(A B C D E) -> '(A D C B E) -> '(B 'before E)
	// '(A B C D E) -> '(A C B D E) -> '(B 'after C)

	Context("With PositionBefore used as position", func() {
		existing := asMovable([]string{"A", "B", "C", "D", "E"})
		Context("when doing a direct move with entries reordering", func() {
			It("should put reordered entries directly before pivot point", func() {
				// '(A B C D E) -> '(A D C B E)
				entries := asMovable([]string{"C", "B"})
				moves, err := movement.MoveGroup(
					movement.PositionBefore{Directly: true, Pivot: "E"},
					entries, existing,
				)

				Expect(err).ToNot(HaveOccurred())
				Expect(moves).To(HaveLen(2))

				Expect(moves[0].Movable.EntryName()).To(Equal("C"))
				Expect(moves[0].Where).To(Equal(movement.ActionWhereBefore))
				Expect(moves[0].Destination.EntryName()).To(Equal("E"))

				Expect(moves[1].Movable.EntryName()).To(Equal("B"))
				Expect(moves[1].Where).To(Equal(movement.ActionWhereAfter))
				Expect(moves[1].Destination.EntryName()).To(Equal("C"))
			})
		})
		Context("when doing a non direct move with entries reordering", func() {
			It("should reorder entries in-place without moving them around", func() {
				// '(A B C D E) -> '(A C B D E)
				entries := asMovable([]string{"C", "B"})
				moves, err := movement.MoveGroup(
					movement.PositionBefore{Directly: false, Pivot: "E"},
					entries, existing,
				)

				Expect(err).ToNot(HaveOccurred())
				Expect(moves).To(HaveLen(1))

				Expect(moves[0].Movable.EntryName()).To(Equal("C"))
				Expect(moves[0].Where).To(Equal(movement.ActionWhereAfter))
				Expect(moves[0].Destination.EntryName()).To(Equal("A"))
			})
		})
		Context("when direct position relative to the pivot is not required", func() {
			Context("and moved entries are already before pivot point", func() {
				It("should not generate any move actions", func() {
					// '(A B C D E) -> '(A B C D E)
					entries := asMovable([]string{"A", "B"})
					moves, err := movement.MoveGroup(
						movement.PositionBefore{Directly: false, Pivot: "D"},
						entries, existing,
					)

					Expect(err).ToNot(HaveOccurred())
					Expect(moves).To(HaveLen(0))
				})
			})
			Context("and moved entries are at the end of the list", func() {
				It("should generate correct move actions", func() {
					// '(A B C D E) -> '(A B D E C)
					entries := asMovable([]string{"D", "E"})
					moves, err := movement.MoveGroup(
						movement.PositionBefore{Directly: false, Pivot: "C"},
						entries, existing,
					)

					Expect(err).ToNot(HaveOccurred())
					Expect(moves).To(HaveLen(2))

					Expect(moves[0].Movable.EntryName()).To(Equal("D"))
					Expect(moves[0].Where).To(Equal(movement.ActionWhereAfter))
					Expect(moves[0].Destination.EntryName()).To(Equal("B"))

					Expect(moves[1].Movable.EntryName()).To(Equal("E"))
					Expect(moves[1].Where).To(Equal(movement.ActionWhereAfter))
					Expect(moves[1].Destination.EntryName()).To(Equal("D"))
				})
			})
			Context("and moved entries are out of order", func() {
				It("should generate a single command to move B before D", func() {
					// '(A B C D E) -> '(B A C D E)
					entries := asMovable([]string{"B", "A"})
					moves, err := movement.MoveGroup(
						movement.PositionBefore{Directly: false, Pivot: "D"},
						entries, existing,
					)

					Expect(err).ToNot(HaveOccurred())
					Expect(moves).To(HaveLen(1))

					Expect(moves[0].Movable.EntryName()).To(Equal("B"))
					Expect(moves[0].Where).To(Equal(movement.ActionWhereFirst))
					Expect(moves[0].Destination).To(BeNil())
				})
			})
			Context("and moved entries are out of order", func() {
				Context("and other entries preceed them", func() {
					It("shold generate a single move command", func() {
						// '(A B C D E) -> '(A C B D E)
						entries := asMovable([]string{"C", "B"})
						moves, err := movement.MoveGroup(
							movement.PositionBefore{Directly: false, Pivot: "D"},
							entries, existing,
						)

						Expect(err).ToNot(HaveOccurred())
						Expect(moves).To(HaveLen(1))

						Expect(moves[0].Movable.EntryName()).To(Equal("C"))
						Expect(moves[0].Where).To(Equal(movement.ActionWhereAfter))
						Expect(moves[0].Destination.EntryName()).To(Equal("A"))
					})
				})
				Context("and moved entries are after the pivot", func() {
					Context("and other entries are left after the pivot", func() {
						It("should generate correct movement plan", func() {
							// '(A B C D E) -> '(A D E B C)
							entries := asMovable([]string{"D", "E"})
							moves, err := movement.MoveGroup(
								movement.PositionBefore{Directly: false, Pivot: "B"},
								entries, existing,
							)

							Expect(err).ToNot(HaveOccurred())
							Expect(moves).To(HaveLen(2))

							Expect(moves[0].Movable.EntryName()).To(Equal("D"))
							Expect(moves[0].Where).To(Equal(movement.ActionWhereAfter))
							Expect(moves[0].Destination.EntryName()).To(Equal("A"))

							Expect(moves[1].Movable.EntryName()).To(Equal("E"))
							Expect(moves[1].Where).To(Equal(movement.ActionWhereAfter))
							Expect(moves[1].Destination.EntryName()).To(Equal("D"))
						})
					})
				})
			})
			Context("and moved entries are out of order", func() {
				It("should generate a single command to move B before D", func() {
					// '(A B C D E) -> '(A B C D E)
					entries := asMovable([]string{"A", "C"})
					moves, err := movement.MoveGroup(
						movement.PositionBefore{Directly: false, Pivot: "D"},
						entries, existing,
					)

					Expect(err).ToNot(HaveOccurred())
					Expect(moves).To(HaveLen(0))
				})
			})
			Context("and moved entries are out of order", func() {
				It("should generate a single command to move B before D", func() {
					// '(A B C D E) -> '(A C B D E)
					entries := asMovable([]string{"A", "C", "B"})
					moves, err := movement.MoveGroup(
						movement.PositionBefore{Directly: false, Pivot: "D"},
						entries, existing,
					)

					Expect(err).ToNot(HaveOccurred())
					Expect(moves).To(HaveLen(1))

					Expect(moves[0].Movable.EntryName()).To(Equal("C"))
					Expect(moves[0].Where).To(Equal(movement.ActionWhereAfter))
					Expect(moves[0].Destination.EntryName()).To(Equal("A"))
				})
			})
		})
		Context("when direct position relative to the pivot is required", func() {
			It("should generate required move actions", func() {
				// '(A B C D E) -> '(C A B D E)
				entries := asMovable([]string{"D", "E"})
				moves, err := movement.MoveGroup(
					movement.PositionBefore{Directly: true, Pivot: "C"},
					entries, existing,
				)

				Expect(err).ToNot(HaveOccurred())
				Expect(moves).To(HaveLen(2))

				Expect(moves[0].Movable.EntryName()).To(Equal("D"))
				Expect(moves[0].Where).To(Equal(movement.ActionWhereBefore))
				Expect(moves[0].Destination.EntryName()).To(Equal("C"))

				Expect(moves[1].Movable.EntryName()).To(Equal("E"))
				Expect(moves[1].Where).To(Equal(movement.ActionWhereAfter))
				Expect(moves[1].Destination.EntryName()).To(Equal("D"))
			})
		})
		Context("when passing single Movement to MoveGroups()", func() {
			existing := asMovable([]string{"A", "B", "C", "D", "E"})
			It("should return a set of move actions that describe it", func() {
				// '(A B C D E) -> '(A D B C E)
				entries := asMovable([]string{"B", "C"})
				moves, err := movement.MoveGroup(
					movement.PositionBefore{Directly: true, Pivot: "E"},
					entries, existing)

				Expect(err).ToNot(HaveOccurred())
				Expect(moves).To(HaveLen(2))
			})
		})
	})
})

var _ = Describe("MoveGroups()", Label("MoveGroups"), func() {
	existing := asMovable([]string{"A", "B", "C", "D", "E"})
	Context("when passing single Movement to MoveGroups()", func() {
		It("should return a set of move actions that describe it", func() {
			// '(A B C D E) -> '(A D B C E)
			entries := asMovable([]string{"B", "C"})
			movements := []movement.Movement{{
				Entries: entries,
				Position: movement.PositionBefore{
					Directly: true,
					Pivot:    "E",
				}}}
			moves, err := movement.MoveGroups[Mock](existing, movements)

			Expect(err).ToNot(HaveOccurred())
			Expect(moves).To(HaveLen(2))
		})
	})
	// Context("when passing single Movement to MoveGroups()", func() {
	// 	FIt("should return a set of move actions that describe it", func() {
	// 		// '(A B C D E) -> '(A D B C E) -> '(D B C E A)
	// 		movements := []movement.Movement{
	// 			{
	// 				Entries: asMovable([]string{"B", "C"}),
	// 				Position: movement.PositionBefore{
	// 					Directly: true,
	// 					Pivot:    Mock{"E"}},
	// 			},
	// 			{
	// 				Entries:  asMovable([]string{"A"}),
	// 				Position: movement.PositionLast{},
	// 			},
	// 		}
	// 		moves, err := movement.MoveGroups(existing, movements)

	// 		Expect(err).ToNot(HaveOccurred())
	// 		Expect(moves).To(HaveLen(3))
	// 	})
	// })
})

var _ = Describe("Movement benchmarks", func() {
	BeforeEach(func() {
		if !Label("benchmark").MatchesLabelFilter(GinkgoLabelFilter()) {
			Skip("unless label 'benchmark' is specified.")
		}
	})
	Context("when moving only a few elements", func() {
		It("should generate a simple sequence of actions", Label("benchmark"), func() {
			var elts []string
			elements := 120
			for idx := range elements {
				elts = append(elts, fmt.Sprintf("%d", idx))
			}
			existing := asMovable(elts)

			entries := asMovable([]string{"90", "80", "70", "60", "50", "40"})
			moves, err := movement.MoveGroup(
				movement.PositionBefore{Directly: true, Pivot: "100"},
				entries, existing,
			)

			Expect(err).ToNot(HaveOccurred())
			Expect(moves).To(HaveLen(6))

			Expect(moves[0].Movable.EntryName()).To(Equal("90"))
			Expect(moves[0].Where).To(Equal(movement.ActionWhereBefore))
			Expect(moves[0].Destination.EntryName()).To(Equal("100"))

			Expect(moves[1].Movable.EntryName()).To(Equal("80"))
			Expect(moves[1].Where).To(Equal(movement.ActionWhereAfter))
			Expect(moves[1].Destination.EntryName()).To(Equal("90"))

			Expect(moves[2].Movable.EntryName()).To(Equal("70"))
			Expect(moves[2].Where).To(Equal(movement.ActionWhereAfter))
			Expect(moves[2].Destination.EntryName()).To(Equal("80"))

			Expect(moves[3].Movable.EntryName()).To(Equal("60"))
			Expect(moves[3].Where).To(Equal(movement.ActionWhereAfter))
			Expect(moves[3].Destination.EntryName()).To(Equal("70"))

			Expect(moves[4].Movable.EntryName()).To(Equal("50"))
			Expect(moves[4].Where).To(Equal(movement.ActionWhereAfter))
			Expect(moves[4].Destination.EntryName()).To(Equal("60"))

			Expect(moves[5].Movable.EntryName()).To(Equal("40"))
			Expect(moves[5].Where).To(Equal(movement.ActionWhereAfter))
			Expect(moves[5].Destination.EntryName()).To(Equal("50"))
		})
	})
})
