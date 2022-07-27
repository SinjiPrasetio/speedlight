package main

func doMigrate(arg2, arg3 string) error {
	dsn := getDSN()

	//run migration command
	switch arg2 {
	case "up":
		err := speed.MigrateUp(dsn)
		if err != nil {
			return err
		}
	case "down":
		if arg3 == "all" {
			err := speed.MigrateDownAll(dsn)
			if err != nil {
				return err
			}
		} else {
			err := speed.Steps(-1, dsn)
			if err != nil {
				return err
			}
		}
	case "reset":
		err := speed.MigrateDownAll(dsn)
		if err != nil {
			return err
		}
		err = speed.MigrateUp(dsn)
		if err != nil {
			return err
		}
	default:
		showHelp()
	}
	return nil
}
