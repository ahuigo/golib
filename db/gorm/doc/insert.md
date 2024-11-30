# options
## stmt.Clause[name]
err := db.Clauses(v=clause.OnConflict{ UpdateAll: true, })

    tx:=db.getInstance()    
    tx.Statement.AddClause(v=OnConflict)
        clause.Expression = OnConflict
        tx.stmt['OnConflict']=clause
        
## stmt.Selects
db.Select("xx")

# create
    // finishier_api.go:17
    db.Create().Execute(tx)
        tx = db.getInstance()
        tx.Statement.Dest = value
        tx.callbacks.processors["create"].Execute(tx)

## Execute: f(db=tx)
    // callbacks.go
    func (p *processor) Execute(db *DB=tx) *DB {
        stmt              = db.Statement
        stmt.BuildClauses = p.Clauses // 'insert','values'
        for _, f := range p.fns {
            f(db) //before, create, after, commit
        }

    // prapare
    stmt.Model= data
    stmt.Parse(stmt.Model)
        // statement.go
        // schema/schema.go
        ParseWithSpecialTableName(dest=stmt.Model)
            field = schema.ParseField(fieldStruct) // schema/field.go
                field.DefaultValueInterface=field.DefaultValue 取决于value的是否是常规类型(int/float/string/struct/array/slice),map/不可以
                field.DataType = DataType(dataTyper.GormDataType())
            stmt.Schema.FieldsByDBName["field"] = field
    


## create.go
    // callbacks/create.go
    // build sql
    if db.Statement.SQL.Len() == 0 {
        db.Statement.SQL.Grow(180)
        db.Statement.AddClauseIfNotExists(clause.Insert{})
        db.Statement.AddClause(ConvertToCreateValues(db.Statement))
            ConvertToCreateValues(stmt *gorm.Statement)
                selectColumns= stmt.SelectAndOmitColumns()
                for db in stmt.Schema.DBNames:
                    if field := stmt.Schema.FieldsByDBName[db]; !field.HasDefaultValue || field.DefaultValueInterface != nil {
                    if v, ok := selectColumns[db]; (ok && v) || (!ok && (!restricted ):
                        values.Columns = append(values.Columns, clause.Column{Name: db})

                



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

