package trader

type Strategy interface {
	Run([]Snapshot) error
}
