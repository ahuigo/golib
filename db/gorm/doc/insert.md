    // finishier_api.go:17
    db.Create()
        tx = db.getInstance()
        tx.callbacks.processors["create"].Execute(tx)

    // callbacks.go
    func (p *processor) Execute(db *DB) *DB {
        stmt              = db.Statement
        stmt.BuildClauses = p.Clauses // 'insert','values'
        for _, f := range p.fns {
            f(db) //before, create, after, commit
        }

    // callbacks/create.go
    if db.Statement.SQL.Len() == 0 {
        db.Statement.SQL.Grow(180)
        db.Statement.AddClauseIfNotExists(clause.Insert{})
        db.Statement.AddClause(ConvertToCreateValues(db.Statement))
            ConvertToCreateValues(stmt *gorm.Statement)
               selectColumns= stmt.SelectAndOmitColumns()
                    for stmt.Selects
                        stmt.Schema.FieldsByDBName["value"]

                



        db.Statement.Build(clauses='insert'|'values'|'RETURNING')
    }



    // statement.go
    func (stmt *Statement) Build(clauses ...string):
        for clause in stmt.Clauses:
            clause(stmt)


    //clause/clause.go
    func (c Clause) Build(builder=stmt):
        c.Expression.Build(c.builder)
            1. values.Build(builder)

