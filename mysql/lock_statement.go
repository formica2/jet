package mysql

import "github.com/go-jet/jet/internal/jet"

type LockStatement interface {
	Statement
	READ() Statement
	WRITE() Statement
}

func LOCK(tables ...jet.SerializerTable) LockStatement {
	newLock := &lockStatementImpl{
		Lock:  jet.ClauseStatementBegin{Name: "LOCK TABLES", Tables: tables},
		Read:  jet.ClauseOptional{Name: "READ"},
		Write: jet.ClauseOptional{Name: "WRITE"},
	}

	newLock.StatementImpl = jet.NewStatementImpl(Dialect, jet.LockStatementType, newLock, &newLock.Lock, &newLock.Read, &newLock.Write)

	return newLock
}

type lockStatementImpl struct {
	jet.StatementImpl

	Lock  jet.ClauseStatementBegin
	Read  jet.ClauseOptional
	Write jet.ClauseOptional
}

func (l *lockStatementImpl) READ() Statement {
	l.Read.Show = true
	return l
}

func (l *lockStatementImpl) WRITE() Statement {
	l.Write.Show = true
	return l
}

func UNLOCK_TABLES() Statement {
	newUnlock := &unlockStatementImpl{
		Unlock: jet.ClauseStatementBegin{Name: "UNLOCK TABLES"},
	}

	newUnlock.StatementImpl = jet.NewStatementImpl(Dialect, jet.UnLockStatementType, newUnlock, &newUnlock.Unlock)

	return newUnlock
}

type unlockStatementImpl struct {
	jet.StatementImpl
	Unlock jet.ClauseStatementBegin
}