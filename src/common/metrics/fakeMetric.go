package metrics

type FakeMetric struct{}

func (f FakeMetric) Inc(_ string, _ int) {}
