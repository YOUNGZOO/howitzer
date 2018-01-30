//mongo client 에서 sample data put 용도

//모형 drop all
db.models.drop();

//insert 모형1
db.models.insert(
{
  acn_model:"00001",
  model_desc:"모형예제1",
  timestamp: Date()
});

//insert 모형2
db.models.insert(
{
  acn_model:"00002",
  model_desc:"모형예제2",
  timestamp: Date()
});

//모형 입력검증
db.models.find();
