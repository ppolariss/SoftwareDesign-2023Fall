package log

//type LogFile struct {
//}
//
//func (l *LogFile) Update(command Command) error {
//	if command == nil {
//		return nil
//	}
//	if reflect.TypeOf(command).Elem().Name() != "save" {
//		return nil
//	}
//
//	interval, err := util.GetInterval(util.GetNow(), workspace.CurWorkspace.CreateAt)
//	if err != nil {
//		return err
//	}
//	f, err := os.OpenFile("./logFiles/logFile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
//	if err != nil {
//		return err
//	}
//
//	once.Do(func() {
//		_ = f.Truncate(0)
//		_ = util.Output("session start at "+workspace.CurWorkspace.CreateAt+"\n", f)
//	})
//	_ = util.Output(workspace.CurWorkspace.FileName+" "+interval+"\n", f)
//	return nil
//}
