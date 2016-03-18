package job

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type JobRepository struct {
	C *mgo.Collection
}

func (r *JobRepository) Create(job *Job) error {

	job.Id = bson.NewObjectId()
	job.CreatedAt = time.Now().UTC()
	job.UpdatedAt = time.Now().UTC()
	job.ExpiresAt = time.Now().UTC().Add(60 * 60 * 24 * 7 * 1e9)

	err := r.C.Insert(&job)
	return err
}

func (r *JobRepository) GetAll() []Job {
	var jobs []Job
	iter := r.C.Find(nil).Sort("-createdat").Iter()
	result := Job{}
	for iter.Next(&result) {
		jobs = append(jobs, result)
	}
	return jobs
}

func (r *JobRepository) GetById(id string) (job Job, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&job)
	return
}

func (r *JobRepository) Update(job *Job, userEmail string) error {
	err := r.C.Update(bson.M{"_id": job.Id, "posteremail": userEmail},
		bson.M{"$set": bson.M{
			"description": job.Description,
			"categoryid":  job.CategoryId,
			"company":     job.Company,
			"position":    job.Position,
			"location":    job.Location,
			"updatedat":   time.Now().UTC(),
		}})
	return err

}

func (r *JobRepository) Delete(id string, userEmail string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id), "posteremail": userEmail})
	return err
}

func (r *JobRepository) GetByUser(userEmail string) []Job {
	var result Job
	var jobs []Job
	itx := r.C.Find(bson.M{"posteremail": userEmail}).Iter()
	for itx.Next(&result) {
		jobs = append(jobs, result)

	}
	return jobs

}
