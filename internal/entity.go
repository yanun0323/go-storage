package internal

type File struct {
	ID        string `gorm:"column:id;type:varchar(255);primaryKey"`
	Type      string `gorm:"column:type;type:varchar(255);not null"`
	Content   []byte `gorm:"column:blob;type:longblob;not null"`
	CreatedAt int64  `gorm:"column:created_at;type:bigint;not null;autoCreateTime"`
}

func (File) TableName() string {
	return "file"
}
