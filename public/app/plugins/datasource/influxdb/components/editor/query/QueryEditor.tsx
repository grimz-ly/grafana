import { css } from '@emotion/css';
import React from 'react';

import { QueryEditorProps } from '@grafana/data/src';

import InfluxDatasource from '../../../datasource';
import { buildRawQuery } from '../../../queryUtils';
import { InfluxOptions, InfluxQuery, InfluxVersion } from '../../../types';

import { FluxQueryEditor } from './flux/FluxQueryEditor';
import { QueryEditorModeSwitcher } from './influxql/QueryEditorModeSwitcher';
import { RawInfluxQLEditor } from './influxql/code/RawInfluxQLEditor';
import { VisualInfluxQLEditor as VisualInfluxQLEditor } from './influxql/visual/VisualInfluxQLEditor';
import { SQLQueryEditor } from './sql/SQLQueryEditor';

type Props = QueryEditorProps<InfluxDatasource, InfluxQuery, InfluxOptions>;

export const QueryEditor = ({ query, onChange, onRunQuery, datasource }: Props) => {
  if (datasource.version === InfluxVersion.Flux) {
    return (
      <div className="gf-form-query-content">
        <FluxQueryEditor query={query} onChange={onChange} onRunQuery={onRunQuery} datasource={datasource} />
      </div>
    );
  }

  if (datasource.version === InfluxVersion.SQL) {
    return <SQLQueryEditor query={query} datasource={datasource} onRunQuery={onRunQuery} onChange={onChange} />;
  }

  return (
    <div className={css({ display: 'flex' })}>
      <div className={css({ flexGrow: 1 })}>
        {query.rawQuery ? (
          <RawInfluxQLEditor query={query} onChange={onChange} onRunQuery={onRunQuery} />
        ) : (
          <VisualInfluxQLEditor query={query} onChange={onChange} onRunQuery={onRunQuery} datasource={datasource} />
        )}
      </div>
      <QueryEditorModeSwitcher
        isRaw={query.rawQuery ?? false}
        onChange={(value) => {
          onChange({ ...query, query: buildRawQuery(query), rawQuery: value });
          onRunQuery();
        }}
      />
    </div>
  );
};