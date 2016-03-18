package category

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CategoryRepository struct {
	C *mgo.Collection
}

func (r *CategoryRepository) Create(category *Category) error {
	category.Id = bson.NewObjectId()
	err := r.C.Insert(&category)
	return err
}

func (r *CategoryRepository) Update(category *Category) error {
	err := r.C.Update(bson.M{"_id": category.Id},
		bson.M{"$set": bson.M{
			"name": category.Name,
		}})
	return err

}

func (r *CategoryRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (r *CategoryRepository) GetAll() ([]Category, error) {
	var result []Category
	err := r.C.Find(nil).Sort("name").Iter().All(&result)
	return result, err
	/*var result Category
	for itx.Next(&result) {
		categorys = append(categorys, result)
	}
	return categorys
	*/
}

func (r *CategoryRepository) GetById(id string) (category Category, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&category)
	return

}
