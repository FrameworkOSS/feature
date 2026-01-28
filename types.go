package feature

// Feature is a way to represent a modular logical component which can be inserted into, or even unloaded from, the portal framework at runtime.
type Feature interface {
	API() int            //The major version number of the portal API this feature was built against. Used to prevent loading incompatible features or to translate event I/O between versions at runtime.
	ID() string          //Unique identifier to reference this feature. Must not conflict with other features or the portal will refuse to call Open.
	Name() string        //Friendly display name to easily label this feature.
	Authors() []string   //The creators and contributors who took part in making this feature exist. Strive to personally recognize individual humans!
	Description() string //An explanation of what this feature provides to the portal. Multiple lines are acceptable, but descriptions should be plaintext only. Formatted descriptions should be provided via structured commands that another reader feature can parse.
	Version() string     //A version number (i.e. semver.org) to indicate this feature's current release state and any expected compatibility.

	Open() (err error)                 //So that it can register commands with their arguments and help texts, setup event listeners, log into any services, communicate with other features, etc. If an error is returned, Close is called immediately after.
	Close() (errs []error, retry bool) //Should gracefully close everything possible before returning a combined error trace if anything failed, and if a retry is requested.
	Input(*Event) error                //To tell this feature about any events which are directed to it, like callbacks to its queries or portal broadcasts. If an error is returned, Close is called immediately after.
	Output() (*Event, error)           //To ask this feature if it has an event to distribute via the portal, like callbacks to queries from other features or broadcasts to send out. If an error is returned, an event may still be distributed but Close will be called immediately after.
}

// FeatureBinding maintains a transport between a physical and virtual feature.
type FeatureBinding struct {
	binding     Feature
	id          string
	name        string
	authors     []string
	description string
	version     string
}
