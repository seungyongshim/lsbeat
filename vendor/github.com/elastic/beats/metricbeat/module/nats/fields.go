// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package nats

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "nats", asset.ModuleFieldsPri, AssetNats); err != nil {
		panic(err)
	}
}

// AssetNats returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/nats.
func AssetNats() string {
	return "eJzUmM+S2zYMxu9+CkxO7SH7AD50ptNecugeuul5Q1OwzIlIKADojfP0HUqWLdGSLTdKtqujKeP76QP4B3wPn/GwhmBUVgDqtMI1vHs0Ku9WAAWKZVero7CG31YA0LwJf1ERK1wBMFZoBNdQmhXA1mFVyLp57z0E4/EUOT16qNObTLE+/jISPz2f0p8+gaWgxgUBUaNO1FkB3RmFF2QERlPAlsnD41miT9CnEOQ98oMrTiMdzmc8vBD3f5+ASs/HHR5DwYc/p0TUebyQKYziPI2nNn6KArQFj8rOgmU06e0LUUshoE1DciHat/qG6h+d101+z0FTho1iceToawxz3z15BvqsSmqqwUhH6oJiiZyNXeHtshGi3yAno2xkxqDVAYxVt0ewlcOgcuEYU1Rc0Kwm3tvxibF0oshY5E6c6liNLujP3vCi7sQ6m2FnxopCmQ1sib3RNRSRhxNotnU1sqMC9Dz3nUCs4RdB++sooUf/sDkMi2wW5NifZhAeSz/pEh8giimbtePx949PUDNZFBkFtcQTkAsUWkWls6ZqRRr7+jzAMQhk+Thx1XGUSqypsHjeVmR0wsIa2WLIR+8w0dZx4KBcdbCZqs/jizBcS/ldVjYq1aHdBbC4WNr6RIyepkpv4eVjRKmjcPlEG19EZsobTzFoknfBknehTBuqyUtgZNGAwbSUlNfcGriWphmAVyAnFTuksQl/k+faavF9vJcRO1CK+ZRaKKEUtaT/e0JPkG8koSfe6YRKRS9p3ZLokX/MqpUk4CTRO6BRaLaCUa6dal5P31VpKV537Djtie0+fme1MX55zo9GtxFnYHaojF8iivb7nbFimcDso0Z2o+O3QGfCZsBQOFF2m9g0DBTAU3BKnArwn78/PE3EuPYhkJ/Xv02+dmNu3flZcFHFO6fpnHLE6H/btMvD9uz12RuKe9H3hl+fPEHcCy5xI69P3lDci85E+T7788kTxBR43+NOZcFecRD27fTUxxuHKVP6x2JBvthBltlu29hANbadroALmaGVE51sGvbz+9U7O4YU+79xeaN294O41HkUMK1G6um3FEORDmNgBnzjLaqxO3wQ9238JmKRTktipSCojWOpgW5Er+DsnD7z8KYRfkrznES7ZIFhhA2mlYPTxMU9Hi9qb+FvTaCoD958XdZTb746Hz208dvTXwGbw1wgs88L7XipS3FT5WbP6X72yKbEKZ5/AwAA//9SJQAz"
}
