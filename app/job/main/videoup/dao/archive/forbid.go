package archive

import (
	"context"

	"go-common/app/job/main/videoup/model/archive"
	"go-common/library/database/sql"
	"go-common/library/log"
)

const (
	_inForbidSQL = `INSERT INTO archive_forbid (aid,rank_attr,recommend_attr,dynamic_attr,show_attr) VALUES (?,?,?,?,?) ON DUPLICATE KEY UPDATE 
					rank_attr=?,recommend_attr=?,dynamic_attr=?,show_attr=?`
	_forbidSQL = `SELECT aid,rank_attr,recommend_attr,dynamic_attr,show_attr,on_flow_id FROM archive_forbid WHERE aid=?`
)

// TxUpForbid insert archive_oper.
func (d *Dao) TxUpForbid(tx *sql.Tx, af *archive.ForbidAttr) (rows int64, err error) {
	res, err := tx.Exec(_inForbidSQL, af.Aid, af.RankV, af.RecommendV, af.DynamicV, af.ShowV, af.RankV, af.RecommendV, af.DynamicV, af.ShowV)
	if err != nil {
		log.Error("d.inForbid.Exec error(%v)", err)
		return
	}
	rows, err = res.RowsAffected()
	return
}

// Forbid get a archive forbid.
func (d *Dao) Forbid(c context.Context, aid int64) (af *archive.ForbidAttr, err error) {
	row := d.db.QueryRow(c, _forbidSQL, aid)
	af = &archive.ForbidAttr{Aid: aid}
	if err = row.Scan(&af.Aid, &af.RankV, &af.RecommendV, &af.DynamicV, &af.ShowV, &af.OnFlowID); err != nil {
		if err == sql.ErrNoRows {
			// af = nil // NOTE: for init
			err = nil
		} else {
			log.Error("row.Scan error(%v)", err)
		}
		return
	}
	return
}
